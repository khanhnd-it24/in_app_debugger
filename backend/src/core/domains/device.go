package domains

import (
	"backend/src/common"
	"context"
)

type Device struct {
	DeviceId   string `bson:"device_id"`
	DeviceName string `bson:"device_name"`
}

func NewDevice(deviceId string, deviceName string) *Device {
	return &Device{
		DeviceId:   deviceId,
		DeviceName: deviceName,
	}
}

func (d *Device) SetDeviceId(deviceId string) *Device {
	d.DeviceId = deviceId
	return d
}

func (d *Device) SetDeviceName(deviceName string) *Device {
	d.DeviceName = deviceName
	return d
}

type DeviceRepo interface {
	SaveOnlineDevice(ctx context.Context, Device *Device) (*Device, *common.Error)
	RemoveOnlineDevice(ctx context.Context, deviceId string) *common.Error
	GetDeviceByDeviceId(ctx context.Context, deviceId string) (*Device, *common.Error)
	GetAllDevices(ctx context.Context) ([]*Device, *common.Error)
}

func (d *Device) CollectionName() string {
	return "device"
}
