package services

import (
	"backend/src/common"
	"backend/src/common/log"
	"backend/src/core/domains"
	"backend/src/infra/transport/message"
	"backend/src/infra/transport/mqtt_client"
	"backend/src/present/http/requests"
	"context"
)

type ConsoleService struct {
	consoleRepo domains.ConsoleRepo
	deviceRepo  domains.DeviceRepo
}

func NewConsoleService(
	consoleRepo domains.ConsoleRepo,
	deviceRepo domains.DeviceRepo,
) *ConsoleService {
	return &ConsoleService{
		consoleRepo: consoleRepo,
		deviceRepo:  deviceRepo,
	}
}

func (c *ConsoleService) CreateConsoleUC(ctx context.Context, req *requests.Console) (*domains.Console, *common.Error) {
	device, err := c.deviceRepo.GetDeviceByDeviceId(ctx, req.DeviceId)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	console := domains.NewConsole()
	console.SetDeviceId(device.DeviceId).SetContent(req.Content)

	console, err = c.consoleRepo.Create(ctx, console)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	go func() {
		consoleMsg := message.NewConsoleMsg(console)
		mqtt_client.GlobalClient.Publish(consoleMsg.Topic(), consoleMsg.Payload())
	}()

	return console, nil
}

func (c *ConsoleService) GetConsolesByDeviceId(ctx context.Context, deviceId string) ([]*domains.Console, *common.Error) {
	device, err := c.deviceRepo.GetDeviceByDeviceId(ctx, deviceId)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	consoles, err := c.consoleRepo.GetConsolesByDeviceId(ctx, device.DeviceId)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	return consoles, nil
}
