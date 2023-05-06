package waitlist

import (
	"fmt"

	"github.com/mixarchitecture/i18np"
	"opensource.turistikrota.com/shared/events"
	"opensource.turistikrota.com/shared/helper"
	"opensource.turistikrota.com/waitlist/src/config"
)

type Events interface {
	Joined(e *JoinedEvent)
	Leaved(e *LeavedEvent)
}

type (
	JoinedEvent struct {
		Email      string `json:"email"`
		LeaveToken string `json:"leave_token"`
		Lang       string
	}
	LeavedEvent struct {
		Email string `json:"email"`
		Lang  string
	}
)

type waitlistEvents struct {
	publisher events.Publisher
	topics    config.Topics
	i18n      *i18np.I18n
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
	I18n      *i18np.I18n
}

func NewEvents(cnf EventConfig) Events {
	return &waitlistEvents{
		publisher: cnf.Publisher,
		topics:    cnf.Topics,
		i18n:      cnf.I18n,
	}
}

func (e *waitlistEvents) Joined(event *JoinedEvent) {
	subject := e.i18n.Translate(I18nMessages.WaitlistJoinedMailSubject, event.Lang)
	template := fmt.Sprintf("waitlist/joined.%s", event.Lang)
	_ = e.publisher.Publish(e.topics.Notify.SendMail, helper.Notify.BuildEmail(event.Email, subject, i18np.P{
		"Email": event.Email,
		"Token": event.LeaveToken,
	}, event.Email, template))
}

func (e *waitlistEvents) Leaved(event *LeavedEvent) {
	subject := e.i18n.Translate(I18nMessages.WaitlistLeavedMailSubject, event.Lang)
	template := fmt.Sprintf("waitlist/leaved.%s", event.Lang)
	_ = e.publisher.Publish(e.topics.Notify.SendMail, helper.Notify.BuildEmail(event.Email, subject, i18np.P{
		"Email": event.Email,
	}, event.Email, template))

}
