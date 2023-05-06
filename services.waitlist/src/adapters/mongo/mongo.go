package mongo

import (
	mongo_waitlist "opensource.turistikrota.com/waitlist/src/adapters/mongo/waitlist"
	"opensource.turistikrota.com/waitlist/src/domain/waitlist"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewWaitlist(waitlistFactory waitlist.Factory, collection *mongo.Collection) waitlist.Repository
}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}

func (m *mongodb) NewWaitlist(waitlistFactory waitlist.Factory, collection *mongo.Collection) waitlist.Repository {
	return mongo_waitlist.New(waitlistFactory, collection)
}
