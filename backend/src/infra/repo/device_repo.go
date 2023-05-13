package repo

import (
	"backend/src/common"
	"backend/src/core/domains"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	onlineDeviceFormat = "online.%s"
	onlinePrefix       = "online.*"
)

func NewDeviceRepo(cache redis.UniversalClient) domains.DeviceRepo {
	return &DeviceRepo{
		cache: cache,
	}
}

type DeviceRepo struct {
	cache redis.UniversalClient
}

func (d DeviceRepo) SaveOnlineDevice(ctx context.Context, device *domains.Device) (*domains.Device, *common.Error) {
	onlineDevice := fmt.Sprintf(onlineDeviceFormat, device.DeviceId)
	if err := d.cache.Set(ctx, onlineDevice, device.DeviceName, 60*time.Hour).Err(); err != nil {
		return nil, common.ErrSystemError(ctx, err.Error())
	}
	return device, nil
}

func (d DeviceRepo) RemoveOnlineDevice(ctx context.Context, deviceId string) *common.Error {
	onlineDevice := fmt.Sprintf(onlineDeviceFormat, deviceId)

	isExist := d.cache.Exists(ctx, onlineDevice)
	if isExist.Val() == 0 {
		return nil
	}

	if err := d.cache.Del(ctx, onlineDevice).Err(); err != nil {
		return common.ErrSystemError(ctx, err.Error())
	}
	return nil
}

func (d DeviceRepo) GetDeviceByDeviceId(ctx context.Context, deviceId string) (*domains.Device, *common.Error) {
	device := new(domains.Device)
	device.SetDeviceId(deviceId)
	onlineDevice := fmt.Sprintf(onlineDeviceFormat, deviceId)
	if err := d.cache.Get(ctx, onlineDevice).Scan(&device.DeviceName); err != nil {
		return nil, common.ErrSystemError(ctx, err.Error())
	}
	return device, nil
}

func (d DeviceRepo) GetAllDevices(ctx context.Context) ([]*domains.Device, *common.Error) {
	onlineDevices := make([]*domains.Device, 0)
	iter := d.cache.Scan(ctx, 0, onlinePrefix, 0).Iterator()
	for iter.Next(ctx) {
		var deviceId string
		_, err := fmt.Sscanf(iter.Val(), onlineDeviceFormat, &deviceId)
		if err != nil {
			continue
		}

		device, ierr := d.GetDeviceByDeviceId(ctx, deviceId)
		if ierr != nil {
			continue
		}

		onlineDevices = append(onlineDevices, device)
	}
	if err := iter.Err(); err != nil {
		return nil, common.ErrSystemError(ctx, err.Error())
	}
	return onlineDevices, nil
}
