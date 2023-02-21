package mongo

import (
	"context"

	"github.com/aguspray001/lora-data-abstraction/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDBGPSLoraRepository struct {
	db *mongo.Database
}

type MongoDBGPSLoraRepository interface {
	PostData(ctx context.Context, collectionName string, payload *entity.GPSLora) (*mongo.InsertOneResult, error)
	GetData(ctx context.Context, collectionName string) ([]*entity.GPSLora, error)
}

func NewMongoDBGPSLoraRepository(db *mongo.Database) MongoDBGPSLoraRepository {
	return &mongoDBGPSLoraRepository{db: db}
}
