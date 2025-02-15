<template>
  <div class="container">
    <h1 class="header">My Todo List</h1>

    <div class="input-section">
      <input
          v-model="newTask"
          type="text"
          placeholder="Enter new task..."
          class="task-input"
          @keyup.enter="addTask"
      />
      <button @click="addTask" class="add-button">
        <span class="plus-icon">+</span> Add Task
      </button>
    </div>

    <div class="task-list">
      <div
          v-for="task in tasks"
          :key="task.id"
          class="task-item"
          :class="{ completed: task.completed }"
      >
        <div class="task-content">
          <label class="checkbox-container">
            <input
                type="checkbox"
                :checked="task.completed"
                @change="toggleTask(task.id)"
                class="status-checkbox"
            />
            <span class="checkmark"></span>
          </label>
          <span class="task-text">{{ task.title }}</span>
        </div>
        <button @click="deleteTask(task.id)" class="delete-button">
          🗑️
        </button>
      </div>
    </div>
  </div>
</template>


<script setup lang="ts">
import { ref, onMounted } from 'vue'
import {Task} from "@/vite-env";
import '../src/style.css'

const tasks = ref<Task[]>([])
const newTask = ref('')

onMounted(async () => {
    tasks.value = await window.go.main.App.GetTasks()
})

const addTask = async () => {
    const title = newTask.value.trim()
    if (title) {
        await window.go.main.App.AddTask(title)
        tasks.value = await window.go.main.App.GetTasks()
        newTask.value = ''
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
</script>
