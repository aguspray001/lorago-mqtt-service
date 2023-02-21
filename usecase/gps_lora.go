package usecase

import (
	"context"

	"github.com/aguspray001/lora-data-abstraction/entity"
)

func (u *gpsLoraUsecase) PostData(ctx context.Context, collectionName string, payload *entity.GPSLora) error {
	_, err := u.gpsLoraRepo.PostData(ctx, collectionName, payload)
	if err != nil {
		return err
	}
	return nil
}

func (u *gpsLoraUsecase) GetData(ctx context.Context, collectionName string) ([]*entity.GPSLora, error) {
	result, err := u.gpsLoraRepo.GetData(ctx, collectionName)
	if err != nil {
		return nil, err
	}
	return result, nil
}
