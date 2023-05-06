package waitlist

import (
	"opensource.turistikrota.com/waitlist/src/adapters/mongo/waitlist/entity"
	"opensource.turistikrota.com/waitlist/src/domain/waitlist"
	"context"
	"github.com/mixarchitecture/i18np"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (r *repo) Join(ctx context.Context, e *waitlist.Entity) *i18np.Error {
	exist, _err := r.emailExist(ctx, e.Email)
	if _err != nil {
		return _err
	}
	if exist {
		return r.factory.Errors.JoinEmailAlreadyExists()
	}
	c := &entity.MongoWaitlist{}
	res, err := r.collection.InsertOne(ctx, c.FromEntity(e))
	if err != nil {
		return r.factory.Errors.JoinFailed()

	}
	c.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *repo) emailExist(ctx context.Context, email string) (bool, *i18np.Error) {
	filter := bson.M{
		entity.Fields.Email: email,
	}
	_, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (r *repo) Leave(ctx context.Context, token string) (*waitlist.Entity, *i18np.Error) {
	filter := bson.M{
		entity.Fields.LeaveToken: token,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.LeaveTokenNotFound()
	}
	update := bson.M{
		"$set": bson.M{
			entity.Fields.IsActive:   false,
			entity.Fields.UpdatedAt:  time.Now(),
			entity.Fields.LeaveToken: "",
		},
	}
	_err := r.helper.UpdateOne(ctx, filter, update)
	if _err != nil {
		return nil, _err
	}
	return e.ToEntity(), nil
}
