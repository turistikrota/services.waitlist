package command

import (
	"opensource.turistikrota.com/shared/decorator"
	"opensource.turistikrota.com/waitlist/src/domain/waitlist"
	"context"
	"github.com/mixarchitecture/i18np"
)

type WaitlistJoinCommand struct {
	Email string
	Lang  string
}

type WaitlistJoinResult struct{}

type WaitlistJoinHandler decorator.CommandHandler[WaitlistJoinCommand, *WaitlistJoinResult]

type waitlistJoinHandler struct {
	repo    waitlist.Repository
	factory waitlist.Factory
	events  waitlist.Events
}

type WaitlistJoinHandlerConfig struct {
	Repo     waitlist.Repository
	Factory  waitlist.Factory
	Events   waitlist.Events
	CqrsBase decorator.Base
}

func NewWaitlistJoinHandler(config WaitlistJoinHandlerConfig) WaitlistJoinHandler {
	return decorator.ApplyCommandDecorators[WaitlistJoinCommand, *WaitlistJoinResult](
		waitlistJoinHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h waitlistJoinHandler) Handle(ctx context.Context, command WaitlistJoinCommand) (*WaitlistJoinResult, *i18np.Error) {
	e := h.factory.New(command.Email)
	if err := h.repo.Join(ctx, e); err != nil {
		return nil, err
	}
	h.events.Joined(&waitlist.JoinedEvent{
		Email:      e.Email,
		LeaveToken: e.LeaveToken,
		Lang:       command.Lang,
	})
	return &WaitlistJoinResult{}, nil
}
