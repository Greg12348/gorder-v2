package query

import (
	"context"
	"github.com/Greg12348/gorder-v2/common/genproto/orderpb"
	"github.com/Greg12348/gorder-v2/common/genproto/stockpb"
)

type StockService interface {
	CheckIfItemsInStock(ctx context.Context, items []*orderpb.ItemWithQuantity) (*stockpb.CheckIfItemsInStockResponse, error)
	GetItems(ctx context.Context, itemIDs []string) ([]*orderpb.Item, error)
}
