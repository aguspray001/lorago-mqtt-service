package mqtt

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aguspray001/lora-data-abstraction/entity"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (d *mqttGPSLoraDelivery) MQTTSubs(client mqtt.Client, topic string) error {
	// callback
	var cbSubs mqtt.MessageHandler = func(mqttClient mqtt.Client, msg mqtt.Message) {
		var data *entity.GPSLora
		var err = json.Unmarshal(msg.Payload(), &data)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Printf("id: %d; lng: %f; lat: %f, speed: %f; \n", data.ID, data.Longitude, data.Latitude, data.GPSSpeed)

		// save to mongo db
		err = d.gpsLoraSvc.PostData(context.Background(), fmt.Sprintf("node-%d", data.ID), data)
		if err != nil {
			fmt.Print(err)
		}
	}
	//subs data
	token := client.Subscribe(topic, 0, cbSubs)
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (d *mqttGPSLoraDelivery) MQTTPub(client mqtt.Client, topic string) error {
	return nil
}
