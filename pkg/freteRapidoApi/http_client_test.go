package freteRapidoApi_test

import (
	"context"
	"testeFreteRapido/pkg/freteRapidoApi"
	"testing"
	"time"
)

func TestHTTPClient_Post(t *testing.T) {
	h := freteRapidoApi.NewHTTPClient("https://teste.com", 5*time.Second)

	tests := []struct {
		name    string
		path    string
		headers map[string]string
		body    []byte
		wantErr bool
	}{
		{
			name:    "simple post",
			path:    "/post",
			headers: map[string]string{"Content-Type": "application/json"},
			body:    []byte(`{"test":123}`),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, status, err := h.Post(context.Background(), tt.path, tt.headers, tt.body)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Post() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("status: %d, body: %s", status, string(got))
		})
	}
}
