package devtoys

import (
	"context"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type BaseApp struct {
	ctx context.Context
}

func NewBaseApp() *BaseApp {
	return &BaseApp{}
}

func (a *BaseApp) Startup(ctx context.Context) {
	a.ctx = ctx
}
