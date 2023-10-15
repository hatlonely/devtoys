package devtoys

import (
	"context"

	"github.com/hatlonely/baseconverter"
	"github.com/pkg/errors"
)

type NumberBaseConversionApp struct {
	ctx context.Context
}

func NewNumberBaseConversionApp() *NumberBaseConversionApp {
	return &NumberBaseConversionApp{}
}

func (a *NumberBaseConversionApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

const baseCharactersLower = "0123456789abcdefghijklmnopqrstuvwxyz"
const baseCharactersUpper = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type NumberBaseConversionReq struct {
	Number           string
	InBase           int
	ToBase           int
	LowerCase        bool
	InBaseCharacters string
	ToBaseCharacters string
}

type NumberBaseConversionRes struct {
	Number string
}

func (a *NumberBaseConversionApp) ConvertNumberBase(req *NumberBaseConversionReq) (*NumberBaseConversionRes, error) {
	baseCharacters := baseCharactersUpper
	if req.LowerCase {
		baseCharacters = baseCharactersLower
	}

	if req.InBaseCharacters == "" {
		req.InBaseCharacters = baseCharacters[:req.InBase]
	}
	if req.ToBaseCharacters == "" {
		req.ToBaseCharacters = baseCharacters[:req.ToBase]
	}

	number, err1, err2 := baseconverter.BaseToBase(req.Number, req.InBaseCharacters, req.ToBaseCharacters)
	if err1 != nil {
		return nil, errors.Wrap(err1, "BaseToBase failed")
	}
	if err2 != nil {
		return nil, errors.Wrap(err2, "BaseToBase failed")
	}

	return &NumberBaseConversionRes{
		Number: number,
	}, nil
}
