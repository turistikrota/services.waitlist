package decorator

import (
	"context"

	"github.com/mixarchitecture/i18np"
)

func ApplyQueryDecorators[H any, R any](handler QueryHandler[H, R], base Base) QueryHandler[H, R] {
	return &queryLoggingDecorator[H, R]{
		base: &queryMetricsDecorator[H, R]{
			base:   handler,
			client: base.MetricsClient,
		},
		logger: base.Logger,
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, *i18np.Error)
}
