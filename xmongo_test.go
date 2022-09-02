package xmongo_test

import (
	"context"
	"fmt"
	"time"

	"github.com/ringsaturn/xmongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ExampleFindOne() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database(databaseName).Collection(collectionName)
	res, err := xmongo.FindOne[Record](ctx, collection, bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func ExampleFind() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database(databaseName).Collection(collectionName)
	res, err := xmongo.Find[Record](ctx, collection, bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
