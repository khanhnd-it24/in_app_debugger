package message

import (
	"backend/src/core/domains"
	"encoding/json"
)

type DeviceMsg struct {
	DeviceId   string `json:"deviceId"`
	DeviceName string `json:"deviceName"`
	IsOnline   bool   `json:"isOnline"`
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

func NewDeviceMsg(device *domains.Device, isOnline bool) *DeviceMsg {
	return &DeviceMsg{
		DeviceId:   device.DeviceId,
		DeviceName: device.DeviceName,
		IsOnline:   isOnline,
	}
}
