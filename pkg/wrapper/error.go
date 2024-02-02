package wrapper

import (
	"net/http"
)

var (
	BadRequestErrorCode   = http.StatusBadRequest
	BadGatewayErrorCode   = http.StatusBadGateway
	NotFoundErrorCode     = http.StatusNotFound
	UnauthorizedErrorCode = http.StatusUnauthorized
	ValidationErrorCode   = 4001

	InternalServerErrorCode = http.StatusInternalServerError
)

func BadRequestError(message string, data ...interface{}) CustomError {
	return CustomError{
		Message:        message,
		StatusCode:     BadRequestErrorCode,
		HttpStatusCode: BadRequestErrorCode,
		Data:           data,
	}
}

func InternalServerError(message string, data ...interface{}) CustomError {
	return CustomError{
		Message:        message,
		StatusCode:     InternalServerErrorCode,
		HttpStatusCode: InternalServerErrorCode,
		Data:           data,
	}
}

func ValidationError(message string, data ...interface{}) CustomError {
	return CustomError{
		Message:        message,
		StatusCode:     ValidationErrorCode,
		HttpStatusCode: BadRequestErrorCode,
		Data:           data,
	}
}

func NotFoundError(message string, data ...interface{}) CustomError {
	return CustomError{
		Message:        message,
		StatusCode:     NotFoundErrorCode,
		HttpStatusCode: NotFoundErrorCode,
		Data:           data,
	}
}

func UnauthorizedError(message string, data ...interface{}) CustomError {
	return CustomError{
		Message:        message,
		StatusCode:     UnauthorizedErrorCode,
		HttpStatusCode: UnauthorizedErrorCode,
		Data:           data,
	}
}
