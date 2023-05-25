package message

const (
	MqttDevices = "devices"
	MqttNetwork = "networks/%s"
	MqttConsole = "console/%s"
)

type AbstractMqtt interface {
	Topic() string
	Payload() string
}
