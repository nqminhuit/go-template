package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoClient struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoClient(uri, dbName string) (*MongoClient, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &MongoClient{
		Client: client,
		DB:     client.Database(dbName),
	}, nil
}

func (m *MongoClient) Disconnect(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}
