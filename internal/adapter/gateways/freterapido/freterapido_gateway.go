package freterapido

import (
	"context"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/internal/domain/interfaces"
	"testeFreteRapido/pkg/freteRapidoApi"
)

type Gateway struct {
	client *freteRapidoApi.Client
}

func NewFreteGateway(client *freteRapidoApi.Client) interfaces.QuoteGateway {
	return &Gateway{client: client}
}

func (g *Gateway) Quote(
	ctx context.Context,
	recipientZip int32,
	volumes []entity.VolumeEntity,
) (quotes []entity.QuoteEntity, err error) {

	dispatchers := []freteRapidoApi.DispatcherRequest{
		{
			RegisteredNumber: g.client.RegisteredNumber,
			Zipcode:          DISPATCHER_ZIP_CODE,
			Volumes:          adaptVolumes(volumes),
		},
	}

	recipient := freteRapidoApi.Recipient{Zipcode: int64(recipientZip)}

	resp, err := g.client.QuoteV3(ctx, recipient, dispatchers)
	if err != nil {
		err = adaptError(err)
		return nil, err
	}

	return adaptResponse(resp), nil
}
