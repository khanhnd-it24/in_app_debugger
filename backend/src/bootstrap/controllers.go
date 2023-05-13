package bootstrap

import (
	"backend/src/present/http/controllers"
	"backend/src/present/http/validator"
	"go.uber.org/fx"
)

func BuildControllerModule() fx.Option {
	return fx.Options(
		fx.Provide(controllers.NewBaseController),
		fx.Provide(controllers.NewDeviceController),
		fx.Provide(controllers.NewNetworkController),
		fx.Provide(controllers.NewConsoleController),
	)
}

func BuildValidator() fx.Option {
	return fx.Options(
		fx.Provide(validator.NewValidator),
	)
}
