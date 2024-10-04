package main

import (
	"context"
	"encoding/json"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Squared returns the square of the given number
func (a *App) Squared(n json.Number) int64 {
	num, _ := n.Int64()
	return num * num
}