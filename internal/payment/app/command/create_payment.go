package command

import (
	"github.com/Greg12348/gorder-v2/common/decorator"
	"github.com/Greg12348/gorder-v2/common/genproto/orderpb"
	"github.com/Greg12348/gorder-v2/payment/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type CreatePayment struct {
	Order *orderpb.Order
}

type CreatePaymentHandler decorator.CommandHandler[CreatePayment, string]

type createPaymentHandler struct {
	processor domain.Processor
	orderGPRC OrderService
}

func NewCreatePaymentHandler(
	processor domain.Processor,
	orderGPRC OrderService,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) CreatePaymentHandler {
	return decorator.ApplyCommandDecorators[CreatePayment, string](
		createPaymentHandler{
			processor: processor,
			orderGPRC: orderGPRC,
		},
		logger,
		metricClient)
}

func (c createPaymentHandler) Handle(ctx context.Context, cmd CreatePayment) (string, error) {
	link, err := c.processor.CreatePaymentLink(ctx, cmd.Order)
	if err != nil {
		return "", err
	}
	logrus.Infof("create payment link for order: %s success, payment link: %s", cmd.Order.ID, link)
	newOrder := &orderpb.Order{
		ID:          cmd.Order.ID,
		CustomerID:  cmd.Order.CustomerID,
		Status:      "waiting_for_payment",
		Items:       cmd.Order.Items,
		PaymentLink: link,
	}
	err = c.orderGPRC.UpdateOrder(ctx, newOrder)
	return link, err
}
