package mqtt_client

import (
	"backend/src/common/configs"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type MqttClient struct {
	client mqtt.Client
}

func NewMqttClient() *MqttClient {
	config := configs.Get()
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", config.MQTT.Host))
	opts.SetClientID(config.MQTT.ClientId)
	opts.SetUsername(config.MQTT.Username)
	opts.SetPassword(config.MQTT.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("%v", token.Error())
		panic(token.Error())
	}

	GlobalClient = &MqttClient{
		client: client,
	}

	return GlobalClient
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Printf("Connected to mqtt broker")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Disconnected from mqtt broker")
}

func (m *MqttClient) Disconnect() {
	m.client.Disconnect(250)
}

func (m *MqttClient) Publish(topic, msg string) bool {
	token := m.client.Publish(topic, 0, false, msg)
	return token.Wait()
}

func (m *MqttClient) PublishTimeout(topic, msg string, timeout time.Duration) bool {
	token := m.client.Publish(topic, 0, false, msg)
	return token.WaitTimeout(timeout)
}

func (m *MqttClient) Subscribe(topic string, f func(client mqtt.Client, msg mqtt.Message)) bool {
	token := m.client.Subscribe(topic, 0, f)
	return token.Wait()
}

func (m *MqttClient) SubscribeTimeout(
	topic string,
	f func(client mqtt.Client, msg mqtt.Message),
	timeout time.Duration,
) bool {
	token := m.client.Subscribe(topic, 0, f)
	return token.WaitTimeout(timeout)
}
