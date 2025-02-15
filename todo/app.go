package main

import (
	"context"
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
