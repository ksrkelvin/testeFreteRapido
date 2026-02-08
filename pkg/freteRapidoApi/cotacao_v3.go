package freteRapidoApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) QuoteV3(
	recipient Recipient,
	dispatchers []DispatcherRequest,
) (response ResponseCotacaoFreteV3, err error) {

	req := RequestCotacaoFreteV3{
		Shipper: Shipper{
			RegisteredNumber: c.RegisteredNumber,
			Token:            c.AuthToken,
			PlatformCode:     c.PlatformCode,
		},
		Recipient:      recipient,
		Dispatchers:    dispatchers,
		SimulationType: []int{0},
	}

	payload, _ := json.Marshal(req)

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	body, status, err := c.HTTP.Post("/api/v3/quote/simulate", headers, payload)
	if err != nil {
		return ResponseCotacaoFreteV3{}, err
	}

	if status != http.StatusOK {
		return ResponseCotacaoFreteV3{}, fmt.Errorf("frete rapido error: %s", body)
	}

	err = json.Unmarshal(body, &response)
	return response, err
}
