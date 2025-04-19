package command

import (
	"github.com/Greg12348/gorder-v2/common/genproto/orderpb"
	"golang.org/x/net/context"
)

type OrderService interface {
	UpdateOrder(ctx context.Context, order *orderpb.Order) error
}
