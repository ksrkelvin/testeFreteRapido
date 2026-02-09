package httpx

import "net/http"

func FromResult(data any, err error) (statusCode int, response any) {
	if err == nil {
		if data == nil {
			return http.StatusNoContent, nil
		}
		return http.StatusOK, data
	}

	appErr, ok := err.(*AppError)
	if !ok {
		appErr = Internal(err)
	}

	return appErr.Status, ErrorResponse{
		Code:        appErr.Code,
		Message:     appErr.Message,
		Description: appErr.Description,
	}
}
