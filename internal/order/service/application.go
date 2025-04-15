package service

import (
	"github.com/Greg12348/gorder-v2/common/metrics"
	"github.com/Greg12348/gorder-v2/order/adapters"
	"github.com/Greg12348/gorder-v2/order/app"
	"github.com/Greg12348/gorder-v2/order/app/command"
	"github.com/Greg12348/gorder-v2/order/app/query"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func NewApplication(ctx context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{
			CreateOrder: command.NewCreateOrderHandler(orderRepo, logger, metricClient),
			UpdateOrder: command.NewUpdateOrderHandler(orderRepo, logger, metricClient),
		},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricClient),
		},
	}
}
