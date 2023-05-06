package decorator

import (
	"opensource.turistikrota.com/shared/metrics"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mixarchitecture/i18np"
)

type MetricsClient interface {
	Inc(key string, value int)
}

func NewMetricsClient() MetricsClient {
	return metrics.NoOp{}
}

type commandMetricsDecorator[C any, R any] struct {
	base   CommandHandler[C, R]
	client MetricsClient
}

type queryMetricsDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	client MetricsClient
}

func (d commandMetricsDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err *i18np.Error) {
	start := time.Now()

	actionName := strings.ToLower(generateActionName(cmd))

	defer func() {
		end := time.Since(start)

		d.client.Inc(fmt.Sprintf("commands.%s.duration", actionName), int(end.Milliseconds()))

		if err == nil {
			d.client.Inc(fmt.Sprintf("commands.%s.success", actionName), 1)
		} else {
			d.client.Inc(fmt.Sprintf("commands.%s.failure", actionName), 1)
		}
	}()

	return d.base.Handle(ctx, cmd)
}

func (d queryMetricsDecorator[C, R]) Handle(ctx context.Context, query C) (result R, err *i18np.Error) {
	start := time.Now()

	actionName := strings.ToLower(generateActionName(query))

	defer func() {
		end := time.Since(start)

		d.client.Inc(fmt.Sprintf("querys.%s.duration", actionName), int(end.Seconds()))

		if err == nil {
			d.client.Inc(fmt.Sprintf("querys.%s.success", actionName), 1)
		} else {
			d.client.Inc(fmt.Sprintf("querys.%s.failure", actionName), 1)
		}
	}()

	return d.base.Handle(ctx, query)
}
