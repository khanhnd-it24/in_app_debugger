package requests

type Device struct {
	DeviceId   string `json:"device_id" validate:"required"`
	DeviceName string `json:"device_name" validate:"required"`
	IsOnline   bool   `json:"is_online"`
}
