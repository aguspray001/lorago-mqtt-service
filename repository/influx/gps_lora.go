package influx

import (
	"context"
	"fmt"
	"time"

	"github.com/aguspray001/lora-data-abstraction/entity"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func (r *influxDBGPSLoraRepository) PostData(ctx context.Context, payload *entity.GPSLora) error {
	// save to influx db
	dataLora := influxdb2.NewPointWithMeasurement("stat").
		AddTag("id", fmt.Sprint(payload.ID)).
		AddField("latitude", payload.Latitude).
		AddField("longitude", payload.Longitude).
		AddField("gps_speed", payload.GPSSpeed).
		AddField("transmitter_rssi", payload.TransmitterRSSI).
		AddField("received_rssi", payload.ReceiverRSSI).
		SetTime(time.Now())

	defer r.dbClient.Close()
	err := r.dbWritter.WritePoint(ctx, dataLora)
	if err != nil {
		return err
	}
	return nil
}
func (r *influxDBGPSLoraRepository) GetData(ctx context.Context) ([]*entity.GPSLora, error) {
	return nil, nil
}
