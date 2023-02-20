package config

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTConfig struct {
	Host          string
	Port          string
	Username      string
	Password      string
	ClientID      string
	CleanSession  bool
	AutoReconnect bool
	Retained      bool
	KeepAlive     time.Duration
	MsgChanDept   uint
}

// var msgPubHandler mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
// 	fmt.Printf("pub message: %s from topic: %s\n", m.Payload(), m.Topic())
// }

var CallbackSubs mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	fmt.Printf("callback message: %s from topic: %s\n", m.Payload(), m.Topic())
}

// var connectHandler mqtt.OnConnectHandler = func(c mqtt.Client) {
// 	t := c.Subscribe("topic/test", 0, nil)
// 	t.Wait()
// 	fmt.Println("Connected")
// }

// var connectLostHandler mqtt.ConnectionLostHandler = func(c mqtt.Client, err error) {
// 	fmt.Printf("Connect lost: %v", err)
// }

func (conf *MQTTConfig) NewMQTTConnect() (mqtt.Client, error) {
	broker := conf.Host + ":" + conf.Port
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return nil, token.Error()
	}
	return client, nil
}
