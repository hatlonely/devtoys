package devtoys

import (
	"github.com/hatlonely/baseconverter"
	"github.com/pkg/errors"
)

type NumberBaseConversionApp struct {
	BaseApp
}

func NewNumberBaseConversionApp() *NumberBaseConversionApp {
	return &NumberBaseConversionApp{}
}

const baseCharactersLower = "0123456789abcdefghijklmnopqrstuvwxyz"
const baseCharactersUpper = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type NumberBaseConversionReq struct {
	Number           string
	InBase           int `validate:"omitempty,gte=2,lte=36"`
	ToBase           int `validate:"omitempty,gte=2,lte=36"`
	LowerCase        bool
	InBaseCharacters string
	ToBaseCharacters string
}

type NumberBaseConversionRes struct {
	Number string
}

func (a *NumberBaseConversionApp) ConvertNumberBase(req *NumberBaseConversionReq) (*NumberBaseConversionRes, error) {
	if req.Number == "" {
		return &NumberBaseConversionRes{}, nil
	}

	if err := validate.Struct(req); err != nil {
		return nil, errors.Wrap(err, "validate.Struct failed")
	}

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
