package devtoys

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

type Base64TextReq struct {
	Text    string
	Mode    string `validate:"oneof=encode decode"`
	Type    string `validate:"oneof=std url"`
	Padding bool
}

type Base64TextRes struct {
	Text string
}

func (a *Base64TextApp) Base64Text(req *Base64TextReq) (*Base64TextRes, error) {
	if err := validate.Struct(req); err != nil {
		return nil, errors.Wrap(err, "validate.Struct failed")
	}

	var text string
	var err error
	if req.Mode == "encode" {
		text, err = a.encode(req.Text, req.Type, req.Padding)
		if err != nil {
			return nil, errors.WithMessage(err, "encode failed")
		}
	} else if req.Mode == "decode" {
		text, err = a.decode(req.Text, req.Type, req.Padding)
		if err != nil {
			return nil, errors.WithMessage(err, "decode failed")
		}
	}

	return &Base64TextRes{
		Text: text,
	}, nil
}

func (a *Base64TextApp) encode(text string, type_ string, padding bool) (string, error) {
	if padding && type_ == "std" {
		return base64.StdEncoding.EncodeToString([]byte(text)), nil
	} else if padding && type_ == "url" {
		return base64.URLEncoding.EncodeToString([]byte(text)), nil
	} else if type_ == "std" {
		return base64.RawStdEncoding.EncodeToString([]byte(text)), nil
	} else if type_ == "url" {
		return base64.RawURLEncoding.EncodeToString([]byte(text)), nil
	}
	return "", errors.New("unknown encode type")
}

func (a *Base64TextApp) decode(text string, type_ string, padding bool) (string, error) {
	var err error
	var decoded []byte
	if padding && type_ == "std" {
		decoded, err = base64.StdEncoding.DecodeString(text)
	} else if padding && type_ == "url" {
		decoded, err = base64.URLEncoding.DecodeString(text)
	} else if type_ == "std" {
		decoded, err = base64.RawStdEncoding.DecodeString(text)
	} else if type_ == "url" {
		decoded, err = base64.RawURLEncoding.DecodeString(text)
	} else {
		return "", errors.New("unknown decode type")
	}

	if err != nil {
		return "", errors.Wrap(err, "base64 decode failed")
	}
	return string(decoded), nil
}
