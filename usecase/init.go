package usecase

import (
	"context"

	"github.com/aguspray001/lora-data-abstraction/entity"
	"github.com/aguspray001/lora-data-abstraction/repository/mongo"
)

type gpsLoraUsecase struct {
	gpsLoraRepo mongo.MongoDBGPSLoraRepository
}

type GPSLoraUsecase interface {
	PostData(ctx context.Context, collectionName string, payload *entity.GPSLora) error
	GetData(ctx context.Context, collectionName string) ([]*entity.GPSLora, error)
}

func NewGPSLoraUsecase(repo *mongo.MongoDBGPSLoraRepository) GPSLoraUsecase {
	return &gpsLoraUsecase{gpsLoraRepo: *repo}
}
