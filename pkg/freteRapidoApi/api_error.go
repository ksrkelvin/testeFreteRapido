package freteRapidoApi

import "fmt"

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("frete rapido api error (%d): %s", e.StatusCode, e.Message)
}

var errorByStatus = map[int]string{
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	422: "Unprocessable Entity",
	429: "Too Many Requests",
	500: "Internal Server Error",
	502: "Bad Gateway",
	503: "Service Unavailable",
}

func MapError(statusCode int, apiErr error) *APIError {
	var internalServerError = 500
	if msg, ok := errorByStatus[statusCode]; ok {
		return &APIError{
			StatusCode: statusCode,
			Message:    msg,
		}
	}

	message := errorByStatus[internalServerError]
	if apiErr != nil {
		message = apiErr.Error()
	}

	return &APIError{
		StatusCode: internalServerError,
		Message:    message,
	}
}
