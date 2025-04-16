package query

import (
	"github.com/Greg12348/gorder-v2/common/decorator"
	"github.com/Greg12348/gorder-v2/common/genproto/orderpb"
	domain "github.com/Greg12348/gorder-v2/stock/domain/stock"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type GetItems struct {
	ItemIDs []string
}

type GetItemsHandler decorator.QueryHandler[GetItems, []*orderpb.Item]

type getItemsHandler struct {
	stockPepo domain.Repository
}

func NewGetItemsHandler(
	stockPepo domain.Repository,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) GetItemsHandler {
	if stockPepo == nil {
		panic("nil stockPepo")
	}
	return decorator.ApplyQueryDecorators[GetItems, []*orderpb.Item](
		getItemsHandler{stockPepo: stockPepo},
		logger,
		metricClient,
	)
}

func (g getItemsHandler) Handle(ctx context.Context, query GetItems) ([]*orderpb.Item, error) {
	logrus.Print("getItemsHandler executed")
	items, err := g.stockPepo.GetItem(ctx, query.ItemIDs)
	if err != nil {
		return nil, err
	}
	return items, nil
}
