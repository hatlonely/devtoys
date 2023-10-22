package devtoys

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/hatlonely/go-kit/config"
	"github.com/pkg/errors"
)

type AppOptions struct {
	Theme string
}

type Options struct {
	App AppOptions
}

type App struct {
	options *Options

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

func NewAppWithOptions(options *Options) *App {
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
		options:                 options,
	}
}

func NewAppWithConfig(filename string) (*App, error) {
	cfg, err := config.NewConfigWithSimpleFile(filename)
	if err != nil {
		return nil, err
	}

	var options Options
	if err := cfg.Unmarshal(&options); err != nil {
		return nil, err
	}

	return NewAppWithOptions(&options), nil
}

func NewApp() (*App, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, errors.Wrap(err, "os.UserConfigDir failed")
	}

	fp := filepath.Join(configDir, "devtoys", "setting.json")

	// 如果文件不存在，则创建一个空的文件，文件内容为空的 options
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(fp), 0755); err != nil {
			return nil, errors.Wrap(err, "failed to create config directory")
		}

		f, err := os.Create(fp)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create config file")
		}
		defer f.Close()

		options := &Options{}
		if err := json.NewEncoder(f).Encode(options); err != nil {
			return nil, errors.Wrap(err, "failed to write options to config file")
		}
	}

	return NewAppWithConfig(fp)
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

func (a *App) GetSetting() *Options {
	return a.options
}

func (a *App) SetSetting(options *Options) {
	a.options = options
}
