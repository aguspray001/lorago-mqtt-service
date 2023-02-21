package mongo

import (
	"context"

	"github.com/aguspray001/lora-data-abstraction/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *mongoDBGPSLoraRepository) PostData(ctx context.Context, collectionName string, payload *entity.GPSLora) (*mongo.InsertOneResult, error) {
	result, err := r.db.Collection(collectionName).InsertOne(ctx, payload)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *mongoDBGPSLoraRepository) GetData(ctx context.Context, collectionName string) ([]*entity.GPSLora, error) {
	result := make([]*entity.GPSLora, 0)
	cur, err := r.db.Collection(collectionName).Find(ctx, "")
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	// append data to array
	for cur.Next(ctx) {
		var row *entity.GPSLora
		err := cur.Decode(&row)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}
