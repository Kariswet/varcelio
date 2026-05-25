package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConfig struct {
	Srv        string
	DbName     string
	Collection string
	Ctx        context.Context

	disableFilterStatusArchive bool
}

func NewMongoConfig(dbname, collname string) *MongoConfig {
	return &MongoConfig{
		Srv:        os.Getenv("MONGO_DB_SRV"),
		DbName:     dbname,
		Collection: collname,
		Ctx:        context.Background(),
	}
}

func (mc *MongoConfig) Connect() (*mongo.Client, error) {
	opts := options.Client().ApplyURI(mc.Srv)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (mc *MongoConfig) Disconnect(client *mongo.Client) {
	if client == nil {
		return
	}
}
