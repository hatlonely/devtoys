package devtoys

import (
	"context"
)

type App struct {
	Base64TextApp
	UnixTimestampApp
	NumberBaseConversionApp
	EscapeApp
	TextBinaryConversionApp
	HashApp
	StringConversionApp
	SSLCertificateApp
	PasswordGeneratorApp
}

func NewApp() *App {
	return &App{
		Base64TextApp:           *NewBase64TextApp(),
		UnixTimestampApp:        *NewUnixTimestampApp(),
		NumberBaseConversionApp: *NewNumberBaseConversionApp(),
		EscapeApp:               *NewEscapeApp(),
		TextBinaryConversionApp: *NewTextBinaryConversionApp(),
		HashApp:                 *NewHashApp(),
		StringConversionApp:     *NewStringConversionApp(),
		SSLCertificateApp:       *NewSSLCertificateApp(),
		PasswordGeneratorApp:    *NewPasswordGeneratorApp(),
	}
}

func (a *App) Startup(ctx context.Context) {
	a.Base64TextApp.Startup(ctx)
	a.UnixTimestampApp.Startup(ctx)
	a.NumberBaseConversionApp.Startup(ctx)
	a.EscapeApp.Startup(ctx)
	a.TextBinaryConversionApp.Startup(ctx)
	a.HashApp.Startup(ctx)
	a.StringConversionApp.Startup(ctx)
	a.SSLCertificateApp.Startup(ctx)
	a.PasswordGeneratorApp.Startup(ctx)
}
