package xmongo_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/ringsaturn/xmongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Record struct {
	OID primitive.ObjectID `bson:"_id"`
	Msg string             `bson:"msg"`
}

var repo *xmongo.Repo[Record]

const (
	databaseName   = "foo"
	collectionName = "bar"
)

func init() {
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
	repo, _ = xmongo.NewRepo[Record](client.Database(databaseName).Collection(collectionName))
}

var (
	insertOnce     = &sync.Once{}
	insertManyOnce = &sync.Once{}
)

func TestRepoInsertOne(t *testing.T) {
	insertOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_, err := repo.InsertOne(ctx, Record{OID: primitive.NewObjectID(), Msg: "insert_one"})
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestRepoInsertMany(t *testing.T) {
	insertManyOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_, err := repo.InsertMany(ctx, []Record{
			{OID: primitive.NewObjectID(), Msg: "insert_many_1"},
			{OID: primitive.NewObjectID(), Msg: "insert_many_2"},
			{OID: primitive.NewObjectID(), Msg: "insert_many_3"},
		})
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestRepoFindOne(t *testing.T) {
	TestRepoInsertOne(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := repo.FindOne(ctx, &bson.M{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRepoFind(t *testing.T) {
	TestRepoInsertMany(t)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	c := int64(3)
	_, err := repo.Find(ctx, &bson.M{}, &options.FindOptions{Limit: &c})
	if err != nil {
		t.Fatal(err)
	}
}
