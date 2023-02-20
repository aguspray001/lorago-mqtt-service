package config

import (
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxDBConfig struct {
	Host         string
	Port         string
	Token        string
	Bucket       string
	Organization string
}

func (conf *InfluxDBConfig) NewInfluxDBConnect() (influxdb2.Client, api.WriteAPIBlocking) {
	var influxUrl string = fmt.Sprintf("http://%s:%s", conf.Host, conf.Port)

	client := influxdb2.NewClient(influxUrl, conf.Token)
	writeAPI := client.WriteAPIBlocking(conf.Organization, conf.Bucket)

	return client, writeAPI
}
