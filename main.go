package main

import (
	"changeme/internal/apps"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	base64TextApp := apps.NewBase64TextApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "devtoys",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        base64TextApp.Startup,
		Bind: []interface{}{
			base64TextApp,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
