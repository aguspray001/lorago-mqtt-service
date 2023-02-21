package mqtt

import (
	"github.com/aguspray001/lora-data-abstraction/usecase"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type mqttGPSLoraDelivery struct {
	gpsLoraSvc usecase.GPSLoraUsecase
}

type MqttGPSLoraDelivery interface {
	MQTTSubs(client mqtt.Client, topic string) error
	MQTTPub(client mqtt.Client, topic string) error
}

func NewMQTTGPSLoraDelivery(svc usecase.GPSLoraUsecase) MqttGPSLoraDelivery {
	return &mqttGPSLoraDelivery{gpsLoraSvc: svc}
}
