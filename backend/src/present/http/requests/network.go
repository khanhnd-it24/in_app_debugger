package requests

type Network struct {
	DeviceId string `json:"device_id" validate:"required"`
	Request  string `json:"request" validate:"required"`
	Response string `json:"response" validate:"required"`
}
