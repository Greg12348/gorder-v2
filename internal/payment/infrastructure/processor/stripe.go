package processor

import (
	"context"
	"github.com/Greg12348/gorder-v2/common/genproto/orderpb"
	"github.com/goccy/go-json"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/checkout/session"
)

type StripeProcessor struct {
	apiKey string
}

func NewStripeProcessor(apiKey string) *StripeProcessor {
	if apiKey == "" {
		panic("empty api key")
	}
	stripe.Key = apiKey
	return &StripeProcessor{apiKey: apiKey}
}

var (
	successURL = "http://localhost:8282/success"
)

func (s StripeProcessor) CreatePaymentLink(ct context.Context, order *orderpb.Order) (string, error) {
	var items []*stripe.CheckoutSessionLineItemParams
	for _, item := range order.Items {
		items = append(items, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.PriceID),
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}

	marshalledItems, _ := json.Marshal(order.Items)
	Metadata := map[string]string{
		"orderID":    order.ID,
		"customerID": order.CustomerID,
		"status":     string(order.Status),
		"items":      string(marshalledItems),
	}
	params := &stripe.CheckoutSessionParams{
		Metadata:   Metadata,
		LineItems:  items,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(successURL),
	}
	result, err := session.New(params)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}
