package bootstrap

import (
	"backend/src/infra/transport/mqtt_client"
	"go.uber.org/fx"
)

func BuildMessageQueueModules() fx.Option {
	return fx.Options(
		fx.Invoke(mqtt_client.NewMqttClient),
	)
}
