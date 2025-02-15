package main

import (
	"context"
	"fmt"
)

type App struct {
	ctx context.Context
	tm  *TaskManager
}

func NewApp() *App {
	return &App{
		tm: NewTaskManager(),
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) AddTask(title string, dueDate string, dueTime string, priority string) (Task, error) {
	a.tm.mu.Lock()
	defer a.tm.mu.Unlock()

	if title == "" {
		return Task{}, fmt.Errorf("title cannot be empty")
	}

	newTask := Task{
		ID:        len(a.tm.tasks) + 1,
		Title:     title,
		DueDate:   dueDate,
		DueTime:   dueTime,
		Priority:  priority,
		Completed: false,
	}

	a.tm.tasks = append(a.tm.tasks, newTask)
	if err := a.tm.saveTasks(); err != nil {
		return Task{}, fmt.Errorf("failed to save tasks: %w", err)
	}

	return newTask, nil
}

func (a *App) GetTasks() []Task {
	a.tm.mu.Lock()
	defer a.tm.mu.Unlock()
	return a.tm.tasks
}
