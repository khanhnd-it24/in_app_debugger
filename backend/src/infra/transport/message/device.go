package message

import (
	"backend/src/core/domains"
	"encoding/json"
)

type DeviceMsg struct {
	DeviceId   string `bson:"device_id" json:"deviceId"`
	DeviceName string `bson:"device_name" json:"deviceName"`
}

func (i *DeviceMsg) Topic() string {
	return MqttDevices
}

func (i *DeviceMsg) Payload() string {
	val, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(val)
}

func NewDeviceMsg(device *domains.Device) *DeviceMsg {
	return &DeviceMsg{
		DeviceId:   device.DeviceId,
		DeviceName: device.DeviceName,
	}
}
