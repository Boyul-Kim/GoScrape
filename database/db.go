package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func GetCollection(name string) *mongo.Collection {
	return dbClient.Database(os.Getenv("dbname")).Collection(name)
}

func InitDb() error {
	var err error
	dbClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("db_uri")))
	if err != nil {
		panic(err)
	}

	err = dbClient.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	return nil
}

func DisconnectDb() {
	err := dbClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}
