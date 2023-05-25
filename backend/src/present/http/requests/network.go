package requests

import "backend/src/common/enums"

type Network struct {
	DeviceId   string           `json:"deviceId" validate:"required"`
	Method     enums.HttpMethod `json:"method" validate:"required"`
	Path       string           `json:"path" validate:"required"`
	StatusCode uint             `json:"statusCode" validate:"required"`
	Request    string           `json:"request" validate:"required"`
	Response   string           `json:"response" validate:"required"`
}
