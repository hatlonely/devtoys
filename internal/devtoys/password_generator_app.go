package devtoys

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type PasswordGeneratorApp struct {
	ctx context.Context
}

func NewPasswordGeneratorApp() *PasswordGeneratorApp {
	return &PasswordGeneratorApp{}
}

func (a *PasswordGeneratorApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

type GeneratePasswordReq struct {
	Length         int `validate:"min=1,max=100"`
	EnableDigit    bool
	EnableUpper    bool
	EnableLower    bool
	EnableSpecific bool
	Specific       string
	CharacterSet   string
}

type GeneratePasswordRes struct {
	Text string
}

func (a *PasswordGeneratorApp) GeneratePassword(req *GeneratePasswordReq) (*GeneratePasswordRes, error) {
	if err := validate.Struct(req); err != nil {
		return nil, errors.Wrap(err, "validate.Struct failed")
	}

	charSet := req.CharacterSet
	if charSet == "" {
		if req.EnableDigit {
			charSet += "0123456789"
		}
		if req.EnableUpper {
			charSet += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}
		if req.EnableLower {
			charSet += "abcdefghijklmnopqrstuvwxyz"
		}
		if req.EnableSpecific {
			if req.Specific == "" {
				req.Specific = "!@#$%^&*()_+-="
			}
			charSet += req.Specific
		}
	}
	if charSet == "" {
		return nil, errors.New("character set is required when no option is enabled")
	}

	rand.Seed(time.Now().UnixNano())
	var password strings.Builder
	for i := 0; i < req.Length; i++ {
		index := rand.Intn(len(charSet))
		password.WriteByte(charSet[index])
	}

	return &GeneratePasswordRes{Text: password.String()}, nil
}
