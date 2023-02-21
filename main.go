package main

import (
	"fmt"
	"time"

	"github.com/aguspray001/lora-data-abstraction/config"
	mqttd "github.com/aguspray001/lora-data-abstraction/delivery/mqtt"
	"github.com/aguspray001/lora-data-abstraction/repository/mongo"
	"github.com/aguspray001/lora-data-abstraction/usecase"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MQTTInit() mqtt.Client {
	mqttConf := &config.MQTTConfig{
		Host:          "ws://10.0.2.15",
		Port:          "8885",
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
	// influxConfig := &config.InfluxDBConfig{
	// 	Host:         "10.0.2.15",
	// 	Port:         "8086",
	// 	Bucket:       "lorago_bucket",
	// 	Organization: "lorago_org",
	// 	Token:        "lorago-f2cec5ee-c6c7-4a10-8fd3-9d8cc7e45d45",
	// }

	// mongodb config
	mongoConfig := &config.MongoDB{
		Uri:        "10.0.2.15",
		Port:       "27017",
		DBName:     "gps_data",
		DBUsername: "mongodb_username",
		DBPassword: "mongodb_pass",
	}
	// connect to db
	mongoDB, err := mongoConfig.NewMongoDBConnect()
	if err != nil {
		panic("Cannot connect to DB")
	}
	// init MQTT
	var mqttClient = MQTTInit()
	topic := "testing/lorago"

	// while loop for mqtt
	for {
		gpsLoraRepo := mongo.NewMongoDBGPSLoraRepository(mongoDB)
		gpsLoraUsecase := usecase.NewGPSLoraUsecase(&gpsLoraRepo)
		gpsLoraDelivery := mqttd.NewMQTTGPSLoraDelivery(gpsLoraUsecase)

		err := gpsLoraDelivery.MQTTSubs(mqttClient, topic)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("waktu sekarang %s \n", time.Now())
		time.Sleep(time.Second)
	}
}
