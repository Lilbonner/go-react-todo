interface Task {
    id: number;
    title: string;
    dueDate?: string;
    dueTime?: string;
    priority?: "low" | "medium" | "high";
    completed: boolean;
}

export { Task };