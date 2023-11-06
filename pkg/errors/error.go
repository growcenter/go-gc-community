package errors

import "errors"

type SystemError struct {
	Error   error
	Code    int
	Message string
}

var (
	INTERNAL_SERVER_ERROR = SystemError{
		Error:   errors.New("err_internal_server_error"),
		Code:    500,
		Message: "Internal Server Error",
	}
)