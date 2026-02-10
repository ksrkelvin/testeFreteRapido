package freteRapidoApi

import "time"

type Client struct {
	RegisteredNumber string
	AuthToken        string
	PlatformCode     string
	HTTP             *httpClient
}

func NewClient(authToken string, platformCode string, registeredNumber string) (*Client, error) {
	return &Client{
		RegisteredNumber: registeredNumber,
		AuthToken:        authToken,
		PlatformCode:     platformCode,
		HTTP: newHTTPClient(
			HOST_FRETE_RAPIDO,
			5*time.Second,
		),
	}, nil
}
