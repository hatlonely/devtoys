package devtoys

import (
	"context"
)

type App struct {
	Base64TextApp
	UnixTimestampApp
	NumberBaseConversionApp
}

func NewApp() *App {
	return &App{
		Base64TextApp:           *NewBase64TextApp(),
		UnixTimestampApp:        *NewUnixTimestampApp(),
		NumberBaseConversionApp: *NewNumberBaseConversionApp(),
	}
}

func (a *App) Startup(ctx context.Context) {
	a.Base64TextApp.Startup(ctx)
	a.UnixTimestampApp.Startup(ctx)
	a.NumberBaseConversionApp.Startup(ctx)
}
