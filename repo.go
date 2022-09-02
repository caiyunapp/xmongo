package xmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo[M any] struct {
	collection *mongo.Collection
}

func NewRepo[M any](collection *mongo.Collection) (*Repo[M], error) {
	return &Repo[M]{
		collection: collection,
	}, nil
}

func (r Repo[M]) Find(ctx context.Context, query *bson.M, opts ...*options.FindOptions) ([]M, error) {
	return Find[M](ctx, r.collection, query, opts...)
}

func (r Repo[M]) FindOne(ctx context.Context, query *bson.M, opts ...*options.FindOneOptions) (M, error) {
	return FindOne[M](ctx, r.collection, query, opts...)
}
