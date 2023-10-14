package app

import (
	"context"
)

type App struct {
	Base64TextApp
	UnixTimestampApp
}

func NewApp() *App {
	return &App{
		Base64TextApp:    *NewBase64TextApp(),
		UnixTimestampApp: *NewUnixTimestampApp(),
	}
}

func (a *App) Startup(ctx context.Context) {
	a.Base64TextApp.Startup(ctx)
	a.UnixTimestampApp.Startup(ctx)
}
