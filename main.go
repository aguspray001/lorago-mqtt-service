package main

import (
	// "context"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aguspray001/lora-data-abstraction/config"
	"github.com/aguspray001/lora-data-abstraction/entity"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	// influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func MQTTInit() mqtt.Client {
	mqttConf := &config.MQTTConfig{
		Host: "ws://10.0.2.15",
		Port: "8885",
		// Username:      "agus",
		// Password:      "r4h4s14",
		// ClientID:      "1122800002",
		CleanSession:  true,
		AutoReconnect: true,
		Retained:      true,
		KeepAlive:     3600 * time.Second,
		MsgChanDept:   100,
	}

	client, err := mqttConf.NewMQTTConnect()
	if err != nil {
		panic(err)
	}

	return client
}

func main() {
	influxConfig := &config.InfluxDBConfig{
		Host:         "10.0.2.15",
		Port:         "8086",
		Bucket:       "lorago_bucket",
		Organization: "lorago_org",
		Token:        "lorago-f2cec5ee-c6c7-4a10-8fd3-9d8cc7e45d45",
	}

	influxClient, writter := influxConfig.NewInfluxDBConnect()

	var cbSubs mqtt.MessageHandler = func(mqttClient mqtt.Client, msg mqtt.Message) {
		// fmt.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())
		// var jsonString = `{"id":1,"latitude":10.000000,"longitude":101.000000,"gps_speed":14.43,"transmitter_rssi":-68, "received_rssi":-70}`
		// var jsonData = []byte(jsonString)

		var data entity.GPSLora
		var err = json.Unmarshal(msg.Payload(), &data)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("id: %d; lng: %f; lat: %f, speed: %f; \n", data.ID, data.Longitude, data.Latitude, data.GPSSpeed)

		// save to influx db
		dataLora := influxdb2.NewPointWithMeasurement("stat").
			AddTag("id", fmt.Sprint(data.ID)).
			AddField("latitude", data.Latitude).
			AddField("longitude", data.Longitude).
			AddField("gps_speed", data.GPSSpeed).
			AddField("transmitter_rssi", data.TransmitterRSSI).
			AddField("received_rssi", data.ReceiverRSSI).
			SetTime(time.Now())

		writter.WritePoint(context.Background(), dataLora)
		influxClient.Close()
	}

	for {
		topic := "testing/lorago"
		var mqttClient = MQTTInit()
		mqttClient.Subscribe(topic, 0, cbSubs)
		// fmt.Printf("waktu sekarang %s \n", time.Now())
		time.Sleep(time.Second)
	}
}
