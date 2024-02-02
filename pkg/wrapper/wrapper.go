package wrapper

import "net/http"

type ContextWrapper interface {
	JSON(code int, i interface{}) error
}

type CustomResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

type CustomError struct {
	Message        string      `json:"message"`
	StatusCode     int         `json:"statusCode"`
	HttpStatusCode int         `json:"httpStatusCode"`
	Data           interface{} `json:"data"`
}

func (ce CustomError) Error() string {
	return ce.Message
}

func SendSuccessResponse(ctx ContextWrapper, message string, data interface{}, code int) error {
	response := CustomResponse{
		Success: true,
		Message: message,
		Data:    data,
		Code:    code,
	}

	return ctx.JSON(http.StatusOK, response)
}

func SendErrorResponse(ctx ContextWrapper, err error, data interface{}, code int) error {
	statusCode := code
	httpCode := code
	responseData := data

	if customErr, ok := err.(CustomError); ok {
		statusCode = customErr.StatusCode
		httpCode = customErr.HttpStatusCode
		if customErr.Data != nil || customErr.Data != "" {
			responseData = customErr.Data
		}
	}

	response := CustomResponse{
		Success: false,
		Message: err.Error(),
		Data:    responseData,
		Code:    statusCode,
	}

	return ctx.JSON(httpCode, response)
}
