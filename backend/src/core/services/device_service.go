package services

import (
	"backend/src/common"
	"backend/src/common/log"
	"backend/src/core/domains"
	"backend/src/present/http/requests"
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

func (d *DeviceService) UpsertDeviceUC(ctx context.Context, req *requests.Device) (*domains.Device, *common.Error) {
	device := domains.NewDevice()
	device.SetDeviceId(req.DeviceId).SetDeviceName(req.DeviceName)

	device, err := d.deviceRepo.Upsert(ctx, device)
	if err != nil {
		log.IErr(ctx, err)
		return nil, err
	}

	return device, nil
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
