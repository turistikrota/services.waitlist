package service

import (
	"opensource.turistikrota.com/shared/db/mongo"
	"opensource.turistikrota.com/shared/decorator"
	"opensource.turistikrota.com/shared/events"
	"opensource.turistikrota.com/shared/validator"
	"opensource.turistikrota.com/waitlist/src/adapters"
	"opensource.turistikrota.com/waitlist/src/app"
	"opensource.turistikrota.com/waitlist/src/app/command"
	"opensource.turistikrota.com/waitlist/src/config"
	"opensource.turistikrota.com/waitlist/src/domain/waitlist"
	"github.com/mixarchitecture/i18np"
)

type Config struct {
	App           config.App
	EventEngine   events.Engine
	MongoWaitlist *mongo.DB
	Validator     *validator.Validator
	I18n          *i18np.I18n
}

func NewApplication(cnf Config) app.Application {
	waitlistFactory := waitlist.NewFactory()
	waitlistRepo := adapters.Mongo.NewWaitlist(waitlistFactory, cnf.MongoWaitlist.GetCollection(cnf.App.DB.Waitlist.Collection))
	waitlistEvents := waitlist.NewEvents(waitlist.EventConfig{
		Publisher: cnf.EventEngine,
		Topics:    cnf.App.Topics,
		I18n:      cnf.I18n,
	})

	base := decorator.NewBase()

	return app.Application{
		Commands: app.Commands{
			Leave: command.NewWaitlistLeaveHandler(command.WaitlistLeaveHandlerConfig{
				Repo:     waitlistRepo,
				Events:   waitlistEvents,
				CqrsBase: base,
			}),
			Join: command.NewWaitlistJoinHandler(command.WaitlistJoinHandlerConfig{
				Repo:     waitlistRepo,
				Factory:  waitlistFactory,
				Events:   waitlistEvents,
				CqrsBase: base,
			}),
		},
		Queries: app.Queries{},
	}
}
