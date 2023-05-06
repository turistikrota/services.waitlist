package main

import (
	"opensource.turistikrota.com/shared/db/mongo"
	"opensource.turistikrota.com/shared/env"
	"opensource.turistikrota.com/shared/events/nats"
	"opensource.turistikrota.com/shared/logs"
	"opensource.turistikrota.com/shared/validator"
	"opensource.turistikrota.com/waitlist/src/config"
	"opensource.turistikrota.com/waitlist/src/delivery"
	"opensource.turistikrota.com/waitlist/src/service"
	"context"
	"github.com/mixarchitecture/i18np"
)

func main() {
	logs.Init()
	ctx := context.Background()
	cnf := config.App{}
	env.Load(&cnf)
	i18n := i18np.New(cnf.I18n.Fallback)
	i18n.Load(cnf.I18n.Dir, cnf.I18n.Locales...)
	eventEngine := nats.New(nats.Config{
		Url:     cnf.Nats.Url,
		Streams: cnf.Nats.Streams,
	})
	valid := validator.New(i18n)
	valid.ConnectCustom()
	valid.RegisterTagName()
	m := loadMongo(cnf)
	app := service.NewApplication(service.Config{
		App:           cnf,
		EventEngine:   eventEngine,
		Validator:     valid,
		I18n:          i18n,
		MongoWaitlist: m,
	})
	delivery := delivery.New(delivery.Config{
		App:         app,
		Config:      cnf,
		I18n:        i18n,
		Validator:   valid,
		Ctx:         ctx,
		EventEngine: eventEngine,
	})
	delivery.Load()
}

func loadMongo(config config.App) *mongo.DB {
	uri := mongo.CalcMongoUri(mongo.UriParams{
		Host:  config.DB.Waitlist.Host,
		Port:  config.DB.Waitlist.Port,
		User:  config.DB.Waitlist.Username,
		Pass:  config.DB.Waitlist.Password,
		Db:    config.DB.Waitlist.Database,
		Query: config.DB.Waitlist.Query,
	})
	d, err := mongo.New(uri, config.DB.Waitlist.Database)
	if err != nil {
		panic(err)
	}
	return d
}
