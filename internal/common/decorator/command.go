package decorator

import (
	"context"
	"github.com/sirupsen/logrus"
)

// CommandHandler defines a generic type that receives a Query Q,
// and returns a result R
type CommandHandler[C, R any] interface {
	Handle(ct context.Context, cmd C) (R, error)
}

func ApplyCommandDecorators[C, R any](handler CommandHandler[C, R], logger *logrus.Entry, metricsClient MetricsClient) CommandHandler[C, R] {
	return queryLoggingDecorator[C, R]{
		logger: logger,
		base: queryMetricsDecorator[C, R]{
			base:   handler,
			client: metricsClient,
		},
	}
}
