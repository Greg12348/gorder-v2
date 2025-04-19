package query

import (
	"context"
	"github.com/Greg12348/gorder-v2/common/decorator"
	"github.com/Greg12348/gorder-v2/common/genproto/orderpb"
	domain "github.com/Greg12348/gorder-v2/stock/domain/stock"
	"github.com/sirupsen/logrus"
)

type CheckIfItemsInStock struct {
	Items []*orderpb.ItemWithQuantity
}

type CheckIfItemsInStockHandler decorator.QueryHandler[CheckIfItemsInStock, []*orderpb.Item]

type checkIfItemsInStockHandler struct {
	stockPepo domain.Repository
}

func NewCheckIfItemsInStockHandler(
	stockPepo domain.Repository,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) CheckIfItemsInStockHandler {
	if stockPepo == nil {
		panic("nil stockPepo")
	}
	return decorator.ApplyQueryDecorators[CheckIfItemsInStock, []*orderpb.Item](
		checkIfItemsInStockHandler{stockPepo: stockPepo},
		logger,
		metricClient,
	)
}

var stub = map[string]string{
	"1": "price_1REgV2D1kdOt6iHJ1q4AoAUZ",
	"2": "price_1RFcyDD1kdOt6iHJYbq8pDG0",
}

func (g checkIfItemsInStockHandler) Handle(ctx context.Context, query CheckIfItemsInStock) ([]*orderpb.Item, error) {
	var res []*orderpb.Item
	for _, i := range query.Items {
		// TODO: catch data from database or stripe
		priceID, ok := stub[i.ID]
		if !ok {
			priceID = stub["1"]
		}
		res = append(res, &orderpb.Item{
			ID:       i.ID,
			Quantity: i.Quantity,
			PriceID:  priceID,
		})
	}
	return res, nil
}
