package domain

import (
	"github.com/Greg12348/gorder-v2/common/genproto/orderpb"
	"golang.org/x/net/context"
)

type Processor interface {
	CreatePaymentLink(ctx context.Context, order *orderpb.Order) (string, error)
}
