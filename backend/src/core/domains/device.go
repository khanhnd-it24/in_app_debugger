package domains

import (
	"backend/src/common"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Device struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	DeviceId   string             `bson:"device_id"`
	DeviceName string             `bson:"device_name"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

func NewDevice() *Device {
	return &Device{
		Id:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (d *Device) SetDeviceId(deviceId string) *Device {
	d.DeviceName = deviceId
	return d
}

func (d *Device) SetDeviceName(deviceName string) *Device {
	d.DeviceName = deviceName
	return d
}

type DeviceRepo interface {
	Upsert(ctx context.Context, Device *Device) (*Device, *common.Error)
	GetDeviceByDeviceId(ctx context.Context, deviceId string) (*Device, *common.Error)
}

func (d *Device) CollectionName() string {
	return "device"
}
