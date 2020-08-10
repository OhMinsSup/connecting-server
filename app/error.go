package app

import "net/http"

type ErrorStatus string

type ErrorException struct {
	Code          int
	Message       ErrorStatus
	ResultMessage string
}

const (
	BadRequest    = ErrorStatus("BAD_REQUEST")
	NotFound      = ErrorStatus("NOT_FOUND")
	Forbidden     = ErrorStatus("FORBIDDEN")
	InteralServer = ErrorStatus("INTERNAL_SERVER_ERROR")
	UnAuthorized  = ErrorStatus("UNAUTHORIZED")
)

func UnAuthorizedErrorResponse(msg string) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusUnauthorized,
		Message:       UnAuthorized,
		ResultMessage: msg,
	}
	return &exception
}

func BadRequestErrorResponse(msg string) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusBadRequest,
		Message:       BadRequest,
		ResultMessage: msg,
	}
	return &exception
}

func NotFoundErrorResponse(msg string) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusNotFound,
		Message:       NotFound,
		ResultMessage: msg,
	}
	return &exception
}

func ForbiddenErrorResponse(msg string) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusForbidden,
		Message:       Forbidden,
		ResultMessage: msg,
	}
	return &exception
}

func InteralServerErrorResponse(msg string) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusInternalServerError,
		Message:       InteralServer,
		ResultMessage: msg,
	}
	return &exception
}
