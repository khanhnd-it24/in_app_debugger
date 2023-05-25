package requests

type Console struct {
	DeviceId string `json:"deviceId" validate:"required"`
	Content  string `json:"content" validate:"required"`
}
