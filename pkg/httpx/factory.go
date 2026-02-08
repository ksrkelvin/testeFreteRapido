package httpx

import "net/http"

func NotFound(message, description string) *AppError {
	return &AppError{
		Status:      http.StatusNotFound,
		Code:        "NOT_FOUND",
		Message:     message,
		Description: description,
	}
}

func InvalidInput(message, description string) *AppError {
	return &AppError{
		Status:      http.StatusBadRequest,
		Code:        "INVALID_INPUT",
		Message:     message,
		Description: description,
	}
}

func Conflict(message, description string) *AppError {
	return &AppError{
		Status:      http.StatusConflict,
		Code:        "CONFLICT",
		Message:     message,
		Description: description,
	}
}

func Unauthorized(message, description string) *AppError {
	return &AppError{
		Status:      http.StatusUnauthorized,
		Code:        "UNAUTHORIZED",
		Message:     message,
		Description: description,
	}
}

func Internal(err error) *AppError {
	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_ERROR",
		Message: "internal server error",
		Err:     err,
	}
}
