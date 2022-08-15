package xmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Decode with generic support.
func Decode[T any](ctx context.Context, cursor mongo.Cursor, stopWhenDecodeErr bool) ([]T, error) {
	res := make([]T, 0)
	for cursor.Next(ctx) {
		newT := new(T)
		if err := cursor.Decode(newT); err != nil {
			if stopWhenDecodeErr {
				return nil, err
			}
			continue
		}
		res = append(res, *newT)
	}
	return res, nil
}
