package services

import (
	"backend/src/common"
	"backend/src/common/log"
	"backend/src/core/domains"
	"backend/src/present/http/requests"
	"backend/src/present/http/responses"
	"context"
)

type DeviceService struct {
	deviceRepo domains.DeviceRepo
}

func NewDeviceService(
	deviceRepo domains.DeviceRepo,
) *DeviceService {
	return &DeviceService{
		deviceRepo: deviceRepo,
	}
}

func (d *DeviceService) UpdateDeviceUC(ctx context.Context, req *requests.Device) (*responses.UpdateDeviceSuccess, *common.Error) {
	device := domains.NewDevice(req.DeviceId, req.DeviceName)

	if req.IsOnline {
		_, err := d.deviceRepo.SaveOnlineDevice(ctx, device)
		if err != nil {
			log.IErr(ctx, err)
			return nil, err
		}

		return &responses.UpdateDeviceSuccess{Success: true}, nil
	}

	err := d.deviceRepo.RemoveOnlineDevice(ctx, device.DeviceId)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	return &responses.UpdateDeviceSuccess{Success: true}, nil
}

func (d *DeviceService) GetDeviceByDeviceId(ctx context.Context, deviceId string) (*domains.Device, *common.Error) {
	device, err := d.deviceRepo.GetDeviceByDeviceId(ctx, deviceId)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	return device, nil
}

func (d *DeviceService) GetAllDevices(ctx context.Context) ([]*domains.Device, *common.Error) {
	devices, err := d.deviceRepo.GetAllDevices(ctx)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	return devices, nil
}
