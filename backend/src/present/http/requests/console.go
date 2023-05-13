package requests

type Console struct {
	DeviceId string `json:"device_id" validate:"required"`
	Content  string `json:"content" validate:"required"`
}
