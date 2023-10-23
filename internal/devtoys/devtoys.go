package devtoys

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

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
	HttpClientApp
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
		HttpClientApp:           *NewHttpClientApp(),
		options:                 options,
	}
}

func NewAppWithConfig(filename string) (*App, error) {
	options, err := readSetting(filename)
	if err != nil {
		return nil, errors.WithMessage(err, "readSetting failed")
	}

	return NewAppWithOptions(options), nil
}

func NewApp() (*App, error) {
	filename, err := getAppConfigPath()
	if err != nil {
		return nil, errors.WithMessage(err, "getAppConfigPath failed")
	}

	return NewAppWithConfig(filename)
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
	a.HttpClientApp.Startup(ctx)
}

func (a *App) Shutdown(ctx context.Context) {
	filename, err := getAppConfigPath()
	if err != nil {
		return
	}

	if err := saveSetting(filename, a.options); err != nil {
		return
	}
}

func (a *App) GetSetting() *Options {
	return a.options
}

func (a *App) SetSetting(options *Options) {
	a.options = options
}

func getAppConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", errors.Wrap(err, "os.UserConfigDir failed")
	}

	return filepath.Join(configDir, "devtoys", "setting.json"), nil
}

func readSetting(filename string) (*Options, error) {
	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &Options{}, nil
		}
		return nil, errors.Wrap(err, "failed to open config file")
	}
	defer f.Close()

	var options Options
	if err := json.NewDecoder(f).Decode(&options); err != nil {
		return nil, errors.Wrap(err, "failed to decode config file")
	}

	return &options, nil
}

func saveSetting(filename string, options *Options) error {
	// 如果文件夹不存在，则创建文件夹
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
			return errors.Wrap(err, "failed to create config directory")
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		return errors.Wrap(err, "failed to create config file")
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(options); err != nil {
		return errors.Wrap(err, "failed to write options to config file")
	}

	return nil
}
