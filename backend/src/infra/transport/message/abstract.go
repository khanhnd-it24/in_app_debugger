package message

const (
	MqttDevices = "devices"
	MqttNetwork = "networks/%s"
	MqttConsole = "consoles/%s"
)

type AbstractMqtt interface {
	Topic() string
	Payload() string
}
