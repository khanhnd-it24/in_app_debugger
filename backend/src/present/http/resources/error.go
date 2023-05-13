package resources

import (
	"backend/src/common"
	"backend/src/core/constant"
	"net/http"
)

type ErrorResponse struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Detail     string `json:"detail,omitempty"`
	Source     string `json:"source"`
	HTTPStatus int    `json:"http_status"`
}

func ConvertErrorToResponse(err *common.Error) *ErrorResponse {
	detail := ""
	if !isInternalError(err) || !constant.IsProdEnv() {
		detail = err.Detail
	}
	return &ErrorResponse{
		Code:       int(err.Code),
		Message:    err.Message,
		Detail:     detail,
		Source:     string(err.Source),
		HTTPStatus: err.HTTPStatus,
	}
}

func isInternalError(err *common.Error) bool {
	return err.GetHttpStatus() >= http.StatusInternalServerError
}
