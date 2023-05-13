package bootstrap

import (
	"backend/src/core/services"
	"go.uber.org/fx"
)

func BuildServicesModules() fx.Option {
	return fx.Options(
		fx.Provide(services.NewDeviceService),
		fx.Provide(services.NewConsoleService),
		fx.Provide(services.NewNetworkService),
	)
}
