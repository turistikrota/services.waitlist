package waitlist

import (
	mongo2 "opensource.turistikrota.com/shared/db/mongo"
	"opensource.turistikrota.com/waitlist/src/adapters/mongo/waitlist/entity"
	"opensource.turistikrota.com/waitlist/src/domain/waitlist"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	factory    waitlist.Factory
	collection *mongo.Collection
	helper     mongo2.Helper[entity.MongoWaitlist, *waitlist.Entity]
}

func New(waitlistFactory waitlist.Factory, collection *mongo.Collection) waitlist.Repository {
	validate(waitlistFactory, collection)
	return &repo{
		factory:    waitlistFactory,
		collection: collection,
		helper:     mongo2.NewHelper[entity.MongoWaitlist, *waitlist.Entity](collection, createEntity),
	}
}

func createEntity() *entity.MongoWaitlist {
	return &entity.MongoWaitlist{}
}

func validate(waitlistFactory waitlist.Factory, collection *mongo.Collection) {
	if waitlistFactory.IsZero() {
		panic("waitlistFactory is zero")
	}
	if collection == nil {
		panic("collection is nil")
	}
}
