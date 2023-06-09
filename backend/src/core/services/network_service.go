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

type NetworkService struct {
	networkRepo domains.NetworkRepo
	deviceRepo  domains.DeviceRepo
}

func NewNetworkService(
	networkRepo domains.NetworkRepo,
	deviceRepo domains.DeviceRepo,
) *NetworkService {
	return &NetworkService{
		networkRepo: networkRepo,
		deviceRepo:  deviceRepo,
	}
}

func (n *NetworkService) CreateNetworkUC(ctx context.Context, req *requests.Network) (*domains.Network, *common.Error) {
	device, err := n.deviceRepo.GetDeviceByDeviceId(ctx, req.DeviceId)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	network := domains.NewNetwork()
	network.
		SetDeviceId(device.DeviceId).
		SetRequest(req.Request).
		SetResponse(req.Response).
		SetStatusCode(req.StatusCode).
		SetMethod(req.Method).
		SetPath(req.Path)

	network, err = n.networkRepo.Create(ctx, network)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	go func() {
		networkMsg := message.NewNetworkMsg(network)
		mqtt_client.GlobalClient.Publish(networkMsg.Topic(), networkMsg.Payload())
	}()

	return network, nil
}

func (n *NetworkService) GetNetworksByDeviceId(ctx context.Context, deviceId string) ([]*domains.Network, *common.Error) {
	device, err := n.deviceRepo.GetDeviceByDeviceId(ctx, deviceId)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	networks, err := n.networkRepo.GetNetworksByDeviceId(ctx, device.DeviceId)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	return networks, nil
}
