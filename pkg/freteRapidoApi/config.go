package freteRapidoApi

import (
	"fmt"
	"os"
)

func NewClient() (*Client, error) {
	host := os.Getenv("FRETE_RAPIDO_API_HOST")
	reg := os.Getenv("FRETE_RAPIDO_API_REGISTERED_NUMBER")
	token := os.Getenv("FRETE_RAPIDO_API_AUTH_TOKEN")
	platform := os.Getenv("FRETE_RAPIDO_API_PLATAFORM_CODE")

	if host == "" || reg == "" || token == "" || platform == "" {
		return nil, fmt.Errorf("frete rapido env vars missing")
	}

	return &Client{
		Host:             host,
		RegisteredNumber: reg,
		AuthToken:        token,
		PlatformCode:     platform,
		HTTP:             newHTTPClient(host),
	}, nil
}
