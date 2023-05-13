package router

import (
	"backend/src/common/configs"
	"backend/src/present/http/controllers"
	"backend/src/present/http/middlewares"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/fx"
)

type RoutersIn struct {
	fx.In
	Engine            *gin.Engine
	DeviceController  *controllers.DeviceController
	NetworkController *controllers.NetworkController
	ConsoleController *controllers.ConsoleController
}

func RegisterHandler(engine *gin.Engine) {
	// recovery
	engine.Use(middlewares.Recovery())
	//tracer
	engine.Use(otelgin.Middleware(configs.Get().Server.Name))
	// log middlewares
	engine.Use(middlewares.Log())
}

func RegisterGinRouters(in RoutersIn) {
	in.Engine.Use(cors.AllowAll())

	group := in.Engine.Group(configs.Get().Server.Http.Prefix)
	group.GET("/ping", middlewares.HealthCheckEndpoint)
	// http swagger serve

	registerPublicRouters(group, in)
}

func registerPublicRouters(r *gin.RouterGroup, in RoutersIn) {
	deviceGroup := r.Group("/devices")
	{
		deviceGroup.GET("/:device_id", in.DeviceController.GetDevice)
		deviceGroup.GET("/", in.DeviceController.GetAllDevices)
		deviceGroup.PUT("/", in.DeviceController.UpdateDevice)
	}

	networkGroup := r.Group("/networks")
	{
		networkGroup.GET("/:device_id", in.NetworkController.GetNetworks)
		networkGroup.POST("/", in.NetworkController.CreateNetwork)
	}

	consoleGroup := r.Group("/consoles")
	{
		consoleGroup.GET("/:device_id", in.ConsoleController.GetConsoles)
		networkGroup.POST("/", in.ConsoleController.CreateConsole)
	}
}
