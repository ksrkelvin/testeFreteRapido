package freteRapidoApi_test

import (
	"testeFreteRapido/pkg/freteRapidoApi"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name             string
		authToken        string
		platformCode     string
		registeredNumber string
		wantErr          bool
	}{
		{
			name:             "basic client creation",
			authToken:        "TOKEN",
			platformCode:     "PLATFORM_CODE",
			registeredNumber: "123456789",
			wantErr:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := freteRapidoApi.NewClient(tt.authToken, tt.platformCode, tt.registeredNumber)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got.AuthToken != tt.authToken {
				t.Errorf("AuthToken = %v, want %v", got.AuthToken, tt.authToken)
			}
			if got.PlatformCode != tt.platformCode {
				t.Errorf("PlatformCode = %v, want %v", got.PlatformCode, tt.platformCode)
			}
			if got.RegisteredNumber != tt.registeredNumber {
				t.Errorf("RegisteredNumber = %v, want %v", got.RegisteredNumber, tt.registeredNumber)
			}
			if got.HTTP == nil {
				t.Error("HTTP client is nil")
			}
		})
	}
}

func TestClient_GetRegisteredNumber(t *testing.T) {
	tests := []struct {
		name             string
		authToken        string
		platformCode     string
		registeredNumber string
		want             string
	}{
		{
			name:             "Get registered number",
			authToken:        "TOKEN",
			platformCode:     "PLATFORM_CODE",
			registeredNumber: "123456789",
			want:             "123456789",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := freteRapidoApi.NewClient(tt.authToken, tt.platformCode, tt.registeredNumber)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			got := c.GetRegisteredNumber()
			if got != tt.want {
				t.Errorf("GetRegisteredNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
