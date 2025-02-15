import { Task } from "./vite-env";

declare global {
    interface Window {
            main: {
                App: {
                    AddTask(
                        title: string,
                        dueDate?: string,
                        dueTime?: string,
                        priority?: string
                    ): Promise<{ result: Task | null; error: string | null }>;
                    GetTasks(): Promise<Task[]>;
                    ToggleTask(id: number): Promise<void>;
                    DeleteTask(id: number): Promise<void>;
                };
            };
    }
}