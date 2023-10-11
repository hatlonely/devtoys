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

func (a *Base64TextApp) Base64Encode(text string, type_ string, raw bool) (string, error) {
	if raw && type_ == "std" {
		return base64.RawStdEncoding.EncodeToString([]byte(text)), nil
	} else if raw && type_ == "url" {
		return base64.RawURLEncoding.EncodeToString([]byte(text)), nil
	} else if type_ == "std" {
		return base64.StdEncoding.EncodeToString([]byte(text)), nil
	} else if type_ == "url" {
		return base64.URLEncoding.EncodeToString([]byte(text)), nil
	}
	return "", errors.New("unknown encode type")
}

func (a *Base64TextApp) Base64Decode(text string, type_ string, raw bool) (string, error) {
	var err error
	var decoded []byte
	if raw && type_ == "std" {
		decoded, err = base64.RawStdEncoding.DecodeString(text)
	} else if raw && type_ == "url" {
		decoded, err = base64.RawURLEncoding.DecodeString(text)
	} else if type_ == "std" {
		decoded, err = base64.StdEncoding.DecodeString(text)
	} else if type_ == "url" {
		decoded, err = base64.URLEncoding.DecodeString(text)
	} else {
		return "", errors.New("unknown decode type")
	}

	if err != nil {
		return "", errors.Wrap(err, "base64 decode failed")
	}
	return string(decoded), nil
}
