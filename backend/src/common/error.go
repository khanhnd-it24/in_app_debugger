package common

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CodeResponse int

const (
	ErrorCodeBadRequest   CodeResponse = 400
	ErrorCodeNotFound     CodeResponse = 404
	ErrorCodeUnauthorized CodeResponse = 401
	ErrorCodeSystemError  CodeResponse = 500
	ErrorCodeForbidden    CodeResponse = 403
)

type Source string

const (
	SourceAPIService Source = "API_Service"
)

type Error struct {
	Code       CodeResponse `json:"code"`
	Message    string       `json:"message"`
	Detail     string       `json:"detail"`
	Source     Source       `json:"source"`
	HTTPStatus int          `json:"http_status"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code:[%d], message:[%s], detail:[%s], source:[%s]", e.Code, e.Message, e.Detail, e.Source)
}

func (e *Error) GetHttpStatus() int {
	return e.HTTPStatus
}

func (e *Error) GetCode() CodeResponse {
	return e.Code
}

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) SetHTTPStatus(status int) *Error {
	e.HTTPStatus = status
	return e
}

func (e *Error) SetMessage(msg string) *Error {
	e.Message = msg
	return e
}

func (e *Error) SetDetail(detail string) *Error {
	e.Detail = detail
	return e
}

func (e *Error) GetDetail() string {
	return e.Detail
}

func (e *Error) SetSource(source Source) *Error {
	e.Source = source
	return e
}

func (e *Error) ToJSon() string {
	data, err := json.Marshal(e)
	if err != nil {
		//Todo fix this
		return "marshal error failed"
	}
	return string(data)
}

var (
	// Status 4xx ********

	ErrUnauthorized = func(ctx context.Context) *Error {
		return &Error{
			Code:       ErrorCodeUnauthorized,
			Message:    DefaultUnauthorizedMessage,
			Source:     SourceAPIService,
			HTTPStatus: http.StatusUnauthorized,
		}
	}

	ErrNotFound = func(ctx context.Context, object, status string) *Error {
		return &Error{
			Code:       ErrorCodeNotFound,
			Message:    getMsg(object, status),
			Source:     SourceAPIService,
			HTTPStatus: http.StatusNotFound,
		}
	}

	ErrBadRequest = func(ctx context.Context) *Error {
		return &Error{
			Code:       ErrorCodeBadRequest,
			Message:    DefaultBadRequestMessage,
			HTTPStatus: http.StatusBadRequest,
			Source:     SourceAPIService,
		}
	}

	// Status 5xx *******

	ErrSystemError = func(ctx context.Context, detail string) *Error {
		return &Error{
			Code:       ErrorCodeSystemError,
			Message:    DefaultServerErrorMessage,
			HTTPStatus: http.StatusInternalServerError,
			Source:     SourceAPIService,
			Detail:     detail,
		}
	}

	ErrForbidden = func(ctx context.Context) *Error {
		return &Error{
			Code:       ErrorCodeForbidden,
			Message:    DefaultForbiddenMessage,
			HTTPStatus: http.StatusForbidden,
			Source:     SourceAPIService,
		}
	}
)

func getMsg(object, status string) string {
	return fmt.Sprintf("%s %s", object, status)
}

const (
	DefaultServerErrorMessage  = "Something has gone wrong, please contact admin"
	DefaultBadRequestMessage   = "Invalid requests"
	DefaultUnauthorizedMessage = "Token invalid"
	DefaultForbiddenMessage    = "Forbidden"
)
