package apps

import (
	"context"
	"encoding/base64"

	"github.com/pkg/errors"
)

type Base64TextApp struct {
	ctx context.Context
}

func NewBase64TextApp() *Base64TextApp {
	return &Base64TextApp{}
}

func (a *Base64TextApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *Base64TextApp) Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func (a *Base64TextApp) Base64Decode(text string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", errors.Wrap(err, "base64 decode failed")
	}
	return string(decoded), nil
}
