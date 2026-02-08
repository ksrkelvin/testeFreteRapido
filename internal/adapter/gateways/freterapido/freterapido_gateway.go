package freterapido

import (
	"context"
	"testeFreteRapido/internal/domain/interfaces"
	"testeFreteRapido/pkg/freteRapidoApi"
)

type Gateway struct {
	client *freteRapidoApi.Client
}

func NewFreteGateway(client *freteRapidoApi.Client) interfaces.FreightGateway {
	return &Gateway{client: client}
}

func (g *Gateway) Quote(
	ctx context.Context,
	recipientZip int32,
	volumes []interfaces.Volume,
) ([]interfaces.FreightQuote, error) {

	dispatchers := []freteRapidoApi.DispatcherRequest{
		{
			RegisteredNumber: g.client.RegisteredNumber,
			Zipcode:          recipientZip,
			Volumes:          adaptVolumes(volumes),
		},
	}

	recipient := freteRapidoApi.Recipient{Zipcode: int64(recipientZip)}

	resp, err := g.client.QuoteV3(recipient, dispatchers)
	if err != nil {
		return nil, err
	}

	return adaptResponse(resp), nil
}
