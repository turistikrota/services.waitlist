package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/mixarchitecture/i18np"
)

type CommandHandler[C any, R any] interface {
	Handle(ctx context.Context, cmd C) (R, *i18np.Error)
}

func ApplyCommandDecorators[H any, R any](handler CommandHandler[H, R], base Base) CommandHandler[H, R] {
	return &commandLoggingDecorator[H, R]{
		base: &commandMetricsDecorator[H, R]{
			base:   handler,
			client: base.MetricsClient,
		},
		logger: base.Logger,
	}
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
