package devtoys

import (
	"context"
)

type BaseApp struct {
	ctx context.Context
}

func NewBaseApp() *BaseApp {
	return &BaseApp{}
}

func (a *BaseApp) Startup(ctx context.Context) {
	a.ctx = ctx
}
