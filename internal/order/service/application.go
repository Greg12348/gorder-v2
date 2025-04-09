package service

import (
	"github.com/Greg12348/gorder-v2/order/app"
	"golang.org/x/net/context"
)

func NewApplication(ctx context.Context) app.Application {
	return app.Application{}
}
