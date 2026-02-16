package freteRapidoApi_test

import (
	"context"
	"testeFreteRapido/pkg/freteRapidoApi"
)

type mockHTTPClient struct {
	ResponseBody []byte
	StatusCode   int
	Err          error
}

func (m *mockHTTPClient) Post(ctx context.Context, path string, headers map[string]string, body []byte) ([]byte, int, error) {
	return m.ResponseBody, m.StatusCode, m.Err
}

func newTestClient(resp []byte, status int, err error) *freteRapidoApi.Client {
	return &freteRapidoApi.Client{
		RegisteredNumber: "123456789",
		AuthToken:        "TOKEN",
		PlatformCode:     "PLATFORM_CODE",
		HTTP: &mockHTTPClient{
			ResponseBody: resp,
			StatusCode:   status,
			Err:          err,
		},
	}
}
