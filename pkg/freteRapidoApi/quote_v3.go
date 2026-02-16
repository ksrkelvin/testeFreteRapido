package freteRapidoApi

import (
	"context"
	"encoding/json"
)

func (c *Client) QuoteV3(
	ctx context.Context,
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

	body, status, err := c.HTTP.Post(ctx, QUOTE_V3_PATH, headers, payload)
	if err != nil {
		return ResponseCotacaoFreteV3{}, err
	}

	if status != 200 {
		return ResponseCotacaoFreteV3{}, MapError(status, err)
	}

	err = json.Unmarshal(body, &response)
	return response, err
}
