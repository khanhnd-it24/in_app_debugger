package requests

type Device struct {
	DeviceId   string `json:"deviceId" validate:"required"`
	DeviceName string `json:"deviceName" validate:"required"`
	IsOnline   bool   `json:"isOnline"`
}
