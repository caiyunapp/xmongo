package xmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne[T any](ctx context.Context, coll *mongo.Collection, filter interface{}, opts ...*options.FindOneOptions) (T, error) {
	t := *new(T)
	res := coll.FindOne(ctx, filter, opts...)
	if err := res.Err(); err != nil {
		return *new(T), err
	}
	if err := res.Decode(&t); err != nil {
		return t, err
	}
	return t, nil
}

func Find[T any](ctx context.Context, coll *mongo.Collection, filter interface{}, opts ...*options.FindOptions) ([]T, error) {
	cursor, err := coll.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	return Decode[T](ctx, cursor)
}

// Decode with generic support.
func Decode[T any](ctx context.Context, cursor *mongo.Cursor) ([]T, error) {
	res := make([]T, 0)
	for cursor.Next(ctx) {
		newT := new(T)
		if err := cursor.Decode(newT); err != nil {
			return nil, err
		}
		res = append(res, *newT)
	}
	return res, nil
}
