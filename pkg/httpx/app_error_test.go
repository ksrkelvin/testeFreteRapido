package httpx_test

import (
	"testeFreteRapido/pkg/httpx"
	"testing"
)

func TestAppError_Error(t *testing.T) {
	tests := []struct {
		name        string
		message     string
		description string
		want        string
	}{
		{
			name:        "BadGateway error",
			message:     "bad gateway",
			description: "upstream service failed",
			want:        "bad gateway",
		},
		{
			name:        "NotFound error",
			message:     "not found",
			description: "resource missing",
			want:        "not found",
		},
		{
			name:        "Internal error",
			message:     "internal server error",
			description: "db connection failed",
			want:        "internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &httpx.AppError{
				Message:     tt.message,
				Description: tt.description,
			}
			got := e.Error()
			if got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
