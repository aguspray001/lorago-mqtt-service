package influx

import (
	"context"

	"github.com/aguspray001/lora-data-abstraction/entity"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type influxDBGPSLoraRepository struct {
	dbClient  influxdb2.Client
	dbWritter api.WriteAPIBlocking
}

type InfluxDBGPSLoraRepository interface {
	PostData(ctx context.Context, payload *entity.GPSLora) error
	GetData(ctx context.Context) ([]*entity.GPSLora, error)
}
