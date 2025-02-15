import { Events } from '@wailsio/runtime';

declare global {
    interface Window {
        go: {
            main: {
                App: {
                    GetTasks(): Promise<Task[]>
                    AddTask(title: string): Promise<void>
                    ToggleTask(id: number): Promise<void>
                    DeleteTask(id: number): Promise<void>
                }
            }
        }
    }
}

export interface Task {
    id: number
    title: string
    completed: boolean
}