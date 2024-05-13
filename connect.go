package mongogo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mi mongoInstance

func InitConnection(connectionString, databaseName string) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))

	defer func() {
		client.Disconnect(context.Background())
	}()

	if err != nil {
		return err
	}

	db := client.Database(databaseName)

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		return err
	}

	mi = mongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}
