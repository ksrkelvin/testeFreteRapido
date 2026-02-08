package freteRapidoApi

type Client struct {
	Host             string
	RegisteredNumber string
	AuthToken        string
	PlatformCode     string
	HTTP             *httpClient
}
