package xmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r Repo[M]) Get(ctx context.Context, oid primitive.ObjectID) (M, error) {
	return r.FindOne(ctx, &bson.M{"_id": oid})
}

func (r Repo[M]) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]M, error) {
	return Find[M](ctx, r.collection, filter, opts...)
}

func (r Repo[M]) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (M, error) {
	return FindOne[M](ctx, r.collection, filter, opts...)
}

func (r Repo[M]) InsertOne(ctx context.Context, doc M, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, doc, opts...)
}

func (r Repo[M]) InsertMany(ctx context.Context, docs []M, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	var inputs []interface{} = make([]interface{}, len(docs))
	for i, d := range docs {
		inputs[i] = d
	}
	return r.collection.InsertMany(ctx, inputs, opts...)
}
