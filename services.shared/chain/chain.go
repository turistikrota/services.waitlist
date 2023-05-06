package chain

import (
	"context"

	"github.com/mixarchitecture/i18np"
)

type Handler[Params any, Result any] func(ctx context.Context, params Params) (*Result, *i18np.Error)

type Chain[Params any, Result any] interface {
	Use(handler ...Handler[Params, Result]) Chain[Params, Result]
	Start(ctx context.Context, params Params) (*Result, *i18np.Error)
	StartWithErr(ctx context.Context, params Params, err *i18np.Error) (*Result, *i18np.Error)
}

type chain[Params any, Result any] struct {
	handlers []Handler[Params, Result]
}

func Make[Params any, Result any]() Chain[Params, Result] {
	return &chain[Params, Result]{}
}

func (c *chain[Params, Result]) Use(handler ...Handler[Params, Result]) Chain[Params, Result] {
	c.handlers = append(c.handlers, handler...)
	return c
}

func (c *chain[Params, Result]) Start(ctx context.Context, params Params) (*Result, *i18np.Error) {
	var result *Result
	var err *i18np.Error
	for _, handler := range c.handlers {
		result, err = handler(ctx, params)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

func (c *chain[Params, Result]) StartWithErr(ctx context.Context, params Params, err *i18np.Error) (*Result, *i18np.Error) {
	if err != nil {
		return nil, err
	}
	return c.Start(ctx, params)
}
