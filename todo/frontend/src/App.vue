<template>
  <div class="container">
    <div class="input-section">
      <input
          v-model="newTask.title"
          type="text"
          placeholder="Task title"
          :class="{ 'error': showError }"
          @keyup.enter="addTask"
      />

      <div class="task-properties">
        <input type="date" v-model="newTask.dueDate">
        <input type="time" v-model="newTask.dueTime">

        <select v-model="newTask.priority">
          <option value="low">Low Priority</option>
          <option value="medium">Medium Priority</option>
          <option value="high">High Priority</option>
        </select>
      </div>

      <button @click="addTask" class="add-button">
        Add Task
      </button>

      <div v-if="showError" class="error-message">
        Task title cannot be empty!
      </div>
    </div>
    <div class="task-list">
      <div
          v-for="task in sortedTasks"
          :key="task.id"
          class="task-item"
          :class="[priorityClass(task.priority), { completed: task.completed }]"
      >
        <div class="task-content">
          <input
              type="checkbox"
              :checked="task.completed"
              @change="toggleTask(task.id)"
          />
          <div class="task-info">
            <span class="task-title">{{ task.title }}</span>
            <div class="task-meta">
              <span v-if="task.dueDate" class="due-date">
                📅 {{ formatDate(task.dueDate) }}
                <span v-if="task.dueTime">⏰ {{ task.dueTime }}</span>
              </span>
              <span class="priority-label">{{ task.priority }} priority</span>
            </div>
          </div>
        </div>
        <button @click="deleteTask(task.id)" class="delete-button">🗑️</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import '../src/style.css'
import { ref, onMounted, computed } from 'vue'

interface Task {
  id: number
  title: string
  dueDate: string
  dueTime: string
  priority: 'low' | 'medium' | 'high'
  completed: boolean
}

declare global {
  interface Window {
    go: {
      main: {
        App: {
          GetTasks(): Promise<Task[]>
          AddTask(title: string,
                  dueDate: string,
                  dueTime: string,
                  priority: string):
                  Promise<void>
          ToggleTask(id: number): Promise<void>
          DeleteTask(id: number): Promise<void>
        }
      }
    }
  }
}

const tasks = ref<Task[]>([])
const newTask = ref({
  title: '',
  dueDate: '',
  dueTime: '',
  priority: 'medium' as 'low' | 'medium' | 'high'
})
const showError = ref(false)

const addTask = async () => {
  if (!newTask.value.title.trim()) {
    showError.value = true
    return
  }

  try {
    await window.go.main.App.AddTask(
        newTask.value.title,
        newTask.value.dueDate,
        newTask.value.dueTime,
        newTask.value.priority
    )

    tasks.value = await window.go.main.App.GetTasks()
    resetForm()
  } catch (error) {
    console.error('Error adding task:', error)
  }
}

const resetForm = () => {
  newTask.value = {
    title: '',
    dueDate: '',
    dueTime: '',
    priority: 'medium'
  }
  showError.value = false
}

const formatDate = (dateString: string) => {
  const options: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }
  return new Date(dateString).toLocaleDateString('en-US', options)
}

const priorityOrder = {
  high: 3,
  medium: 2,
  low: 1
} as const

const sortedTasks = computed(() => {
  return [...tasks.value].sort((a, b) => {
    return priorityOrder[b.priority] - priorityOrder[a.priority]
  })
})

const priorityClass = (priority: 'low' | 'medium' | 'high') => {
  return {
    'low-priority': priority === 'low',
    'medium-priority': priority === 'medium',
    'high-priority': priority === 'high'
  }
}

const toggleTask = async (id: number) => {
  await window.go.main.App.ToggleTask(id)
  tasks.value = await window.go.main.App.GetTasks()
}

const deleteTask = async (id: number) => {
  await window.go.main.App.DeleteTask(id)
  tasks.value = await window.go.main.App.GetTasks()
}

onMounted(async () => {
  tasks.value = await window.go.main.App.GetTasks()
})
</script>/