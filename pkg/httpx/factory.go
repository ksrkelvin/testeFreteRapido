package httpx

import (
	"net/http"
)

func NotFound(message, description string) (appError *AppError) {
	return &AppError{
		StatusCode:  http.StatusNotFound,
		Message:     message,
		Description: description,
	}
}

func BadRequest(message, description string) (appError *AppError) {
	return &AppError{
		StatusCode:  http.StatusBadRequest,
		Message:     message,
		Description: description,
	}
}

func Conflict(message, description string) (appError *AppError) {
	return &AppError{
		StatusCode:  http.StatusConflict,
		Message:     message,
		Description: description,
	}
}

func Internal(err error) (appError *AppError) {
	return &AppError{
		StatusCode:  http.StatusInternalServerError,
		Message:     "internal server error",
		Description: err.Error(),
	}
}

func ServiceUnavailable(message, description string) *AppError {
	return &AppError{
		StatusCode:  http.StatusServiceUnavailable,
		Message:     message,
		Description: description,
	}
}

func BadGateway(message, description string) *AppError {
	return &AppError{
		StatusCode:  http.StatusBadGateway,
		Message:     message,
		Description: description,
	}
}
