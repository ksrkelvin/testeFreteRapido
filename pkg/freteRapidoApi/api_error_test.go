package freteRapidoApi_test

import (
	"fmt"
	"testeFreteRapido/pkg/freteRapidoApi"
	"testing"
)

func TestMapError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		apiErr     error
		want       *freteRapidoApi.APIError
	}{
		{
			name:       "Known status 400",
			statusCode: 400,
			apiErr:     nil,
			want: &freteRapidoApi.APIError{
				StatusCode: 400,
				Message:    "Bad Request",
			},
		},
		{
			name:       "Known status 401",
			statusCode: 401,
			apiErr:     nil,
			want: &freteRapidoApi.APIError{
				StatusCode: 401,
				Message:    "Unauthorized",
			},
		},
		{
			name:       "Unknown status with apiErr",
			statusCode: 999,
			apiErr:     fmt.Errorf("Custom API error"),
			want: &freteRapidoApi.APIError{
				StatusCode: 500,
				Message:    "Custom API error",
			},
		},
		{
			name:       "Unknown status without apiErr",
			statusCode: 999,
			apiErr:     nil,
			want: &freteRapidoApi.APIError{
				StatusCode: 500,
				Message:    "Internal Server Error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := freteRapidoApi.MapError(tt.statusCode, tt.apiErr)

			if got.StatusCode != tt.want.StatusCode {
				t.Errorf("StatusCode = %d, want %d", got.StatusCode, tt.want.StatusCode)
			}
			if got.Message != tt.want.Message {
				t.Errorf("Message = %q, want %q", got.Message, tt.want.Message)
			}
		})
	}
}
