package command

import (
	"opensource.turistikrota.com/shared/decorator"
	"opensource.turistikrota.com/waitlist/src/domain/waitlist"
	"context"
	"github.com/mixarchitecture/i18np"
)

type WaitlistLeaveCommand struct {
	LeaveToken string
	Lang       string
}

type WaitlistLeaveResult struct{}

type WaitlistLeaveHandler decorator.CommandHandler[WaitlistLeaveCommand, *WaitlistLeaveResult]

type waitlistLeaveHandler struct {
	repo   waitlist.Repository
	events waitlist.Events
}

type WaitlistLeaveHandlerConfig struct {
	Repo     waitlist.Repository
	Events   waitlist.Events
	CqrsBase decorator.Base
}

func NewWaitlistLeaveHandler(config WaitlistLeaveHandlerConfig) WaitlistLeaveHandler {
	return decorator.ApplyCommandDecorators[WaitlistLeaveCommand, *WaitlistLeaveResult](
		waitlistLeaveHandler{
			repo:   config.Repo,
			events: config.Events,
		},
		config.CqrsBase,
	)
}

func (h waitlistLeaveHandler) Handle(ctx context.Context, command WaitlistLeaveCommand) (*WaitlistLeaveResult, *i18np.Error) {
	res, err := h.repo.Leave(ctx, command.LeaveToken)
	if err != nil {
		return nil, err
	}
	h.events.Leaved(&waitlist.LeavedEvent{
		Email: res.Email,
		Lang:  command.Lang,
	})
	return &WaitlistLeaveResult{}, nil
}
