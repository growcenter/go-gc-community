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
	DATA_EXIST = SystemError{
		Error:   errors.New("data_exist"),
		Code:    400,
		Message: "Data already exist",
	}
	DATA_INVALID = SystemError{
		Error:   errors.New("data_empty_invalid"),
		Code:    422,
		Message: "The required field on the body request is empty or invalid.",
	}
	UNAUTHORIZED = SystemError{
		Error:   errors.New("Unauthorized"),
		Code:    422,
		Message: "The required field on the body request is empty or invalid.",
	}
)