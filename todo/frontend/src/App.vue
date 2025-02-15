<template>
    <div class="container">
        <h1>Todo List</h1>

        <div class="input-group">
            <input
                v-model="newTask"
                type="text"
                placeholder="New task..."
                @keyup.enter="addTask"
            />
            <button @click="addTask">Add</button>
        </div>

        <ul v-if="tasks.length" class="task-list">
            <li
                v-for="task in tasks"
                :key="task.id"
                :class="{ completed: task.completed }"
            >
                <div class="task-content">
                    <input
                        type="checkbox"
                        :checked="task.completed"
                        @change="toggleTask(task.id)"
                    />
                    <span>{{ task.title }}</span>
                </div>
                <button @click="deleteTask(task.id)" class="delete-btn">×</button>
            </li>
        </ul>

        <p v-else class="empty-state">No tasks yet. Add your first task!</p>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import {Task} from "@/vite-env";


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
