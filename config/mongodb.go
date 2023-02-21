package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Uri        string
	Port       string
	DBName     string
	DBUsername string
	DBPassword string
}

func (m *MongoDB) NewMongoDBConnect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/?authMechanism=SCRAM-SHA-256", m.DBUsername, m.DBPassword, m.Uri, m.Port))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return client.Database(m.DBName), nil
}
