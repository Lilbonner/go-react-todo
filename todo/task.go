package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	DueDate   string `json:"dueDate"`
	DueTime   string `json:"dueTime"`
	Priority  string `json:"priority"`
	Completed bool   `json:"completed"`
}

type TaskManager struct {
	tasks []Task
	mu    sync.Mutex
	file  string
}

func NewTaskManager() *TaskManager {
	configDir, _ := os.UserConfigDir()
	filePath := filepath.Join(configDir, "todolist", "tasks.json")
	os.MkdirAll(filepath.Dir(filePath), 0755)

	tm := &TaskManager{file: filePath}
	tm.loadTasks()
	return tm
}

func (tm *TaskManager) loadTasks() error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	data, err := os.ReadFile(tm.file)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &tm.tasks)
}

func (tm *TaskManager) saveTasks() error {
	data, err := json.MarshalIndent(tm.tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(tm.file, data, 0644)
}
