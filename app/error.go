package app

import "net/http"

type ErrorStatus string

type ErrorException struct {
	Code          int         `json:"code"`
	Message       ErrorStatus `json:"message"`
	ResultCode    int         `json:"result_code"`
	ResultMessage error       `json:"result_message"`
}

const (
	ALREADY_EXIST = ErrorStatus("ALREADY_EXIST")
	BadRequest    = ErrorStatus("BAD_REQUEST")
	NotFound      = ErrorStatus("NOT_FOUND")
	Forbidden     = ErrorStatus("FORBIDDEN")
	InteralServer = ErrorStatus("INTERNAL_SERVER_ERROR")
	UnAuthorized  = ErrorStatus("UNAUTHORIZED")
)

func AlreadyExistsErrorResponse(msg error) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusOK,
		Message:       ALREADY_EXIST,
		ResultCode:    2003,
		ResultMessage: msg,
	}
	return &exception
}

func UnAuthorizedErrorResponse(msg error) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusUnauthorized,
		Message:       UnAuthorized,
		ResultCode:    -1,
		ResultMessage: msg,
	}
	return &exception
}

func BadRequestErrorResponse(msg error) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusBadRequest,
		Message:       BadRequest,
		ResultCode:    -1,
		ResultMessage: msg,
	}
	return &exception
}

func NotFoundErrorResponse(msg error) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusNotFound,
		Message:       NotFound,
		ResultCode:    -1,
		ResultMessage: msg,
	}
	return &exception
}

func ForbiddenErrorResponse(msg error) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusForbidden,
		Message:       Forbidden,
		ResultCode:    -1,
		ResultMessage: msg,
	}
	return &exception
}

func InteralServerErrorResponse(msg error) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusInternalServerError,
		Message:       InteralServer,
		ResultCode:    -1,
		ResultMessage: msg,
	}
	return &exception
}
