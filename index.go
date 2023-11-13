package xmongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// doc: https://www.mongodb.com/docs/v7.0/tutorial/manage-indexes/#minimize-performance-impact-with-a-temporary-index
func ModifyIndex(ctx context.Context, collection *mongo.Collection, name string, newIndex mongo.IndexModel) error {
	// 0. find the existing index
	indexes, err := collection.Indexes().ListSpecifications(ctx, options.ListIndexes())
	if err != nil {
		return fmt.Errorf("list indexes: %w", err)
	}

	var existingIndex *mongo.IndexSpecification
	for _, i := range indexes {
		if i.Name == name {
			existingIndex = i
		}
	}

	if existingIndex == nil {
		return fmt.Errorf("index %s not found", name)
	}

	rawkeys, err := existingIndex.KeysDocument.Elements()
	if err != nil {
		return fmt.Errorf("keys document: %w", err)
	}

	var keys []bson.E
	for _, key := range rawkeys {
		keys = append(keys, bson.E{Key: key.Key(), Value: 1})
	}

	const dummyField = "_tmp_dummy_field"
	keys = append(keys, bson.E{Key: dummyField, Value: 1})

	tmpIndexName := fmt.Sprintf("temp_index_%s", name)

	// 1. create a temporary index
	if _, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: keys,
		Options: &options.IndexOptions{
			Name: &tmpIndexName,
		},
	}); err != nil {
		return fmt.Errorf("create temp index: %w", err)
	}

	// 2. drop the existing index
	if _, err := collection.Indexes().DropOne(ctx, name); err != nil {
		return fmt.Errorf("drop index %s: %w", name, err)
	}

	// 3. create the new index
	if _, err := collection.Indexes().CreateOne(ctx, newIndex); err != nil {
		return fmt.Errorf("create new index: %w", err)
	}

	// 4. drop the temporary index
	if _, err := collection.Indexes().DropOne(ctx, tmpIndexName); err != nil {
		return fmt.Errorf("drop temp index: %w", err)
	}

	return nil
}
