package httpx_test

import (
	"errors"
	"net/http"
	"testing"

	"testeFreteRapido/pkg/httpx"
)

func TestAppErrors(t *testing.T) {
	tests := []struct {
		name        string
		fn          func() *httpx.AppError
		wantStatus  int
		wantMessage string
		wantDesc    string
	}{
		{
			name: "NotFound",
			fn: func() *httpx.AppError {
				return httpx.NotFound("not found", "resource missing")
			},
			wantStatus:  http.StatusNotFound,
			wantMessage: "not found",
			wantDesc:    "resource missing",
		},
		{
			name: "BadRequest",
			fn: func() *httpx.AppError {
				return httpx.BadRequest("bad request", "invalid input")
			},
			wantStatus:  http.StatusBadRequest,
			wantMessage: "bad request",
			wantDesc:    "invalid input",
		},
		{
			name: "Conflict",
			fn: func() *httpx.AppError {
				return httpx.Conflict("conflict", "already exists")
			},
			wantStatus:  http.StatusConflict,
			wantMessage: "conflict",
			wantDesc:    "already exists",
		},
		{
			name: "Internal",
			fn: func() *httpx.AppError {
				return httpx.Internal(errors.New("db connection failed"))
			},
			wantStatus:  http.StatusInternalServerError,
			wantMessage: "internal server error",
			wantDesc:    "db connection failed",
		},
		{
			name: "ServiceUnavailable",
			fn: func() *httpx.AppError {
				return httpx.ServiceUnavailable("service down", "maintenance")
			},
			wantStatus:  http.StatusServiceUnavailable,
			wantMessage: "service down",
			wantDesc:    "maintenance",
		},
		{
			name: "BadGateway",
			fn: func() *httpx.AppError {
				return httpx.BadGateway("bad gateway", "upstream error")
			},
			wantStatus:  http.StatusBadGateway,
			wantMessage: "bad gateway",
			wantDesc:    "upstream error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn()
			if got.StatusCode != tt.wantStatus {
				t.Errorf("StatusCode = %v, want %v", got.StatusCode, tt.wantStatus)
			}
			if got.Message != tt.wantMessage {
				t.Errorf("Message = %v, want %v", got.Message, tt.wantMessage)
			}
			if got.Description != tt.wantDesc {
				t.Errorf("Description = %v, want %v", got.Description, tt.wantDesc)
			}
		})
	}
}
