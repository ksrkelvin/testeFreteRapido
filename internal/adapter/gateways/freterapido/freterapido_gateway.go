package freterapido

import (
	"context"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/internal/domain/interfaces"
	"testeFreteRapido/pkg/freteRapidoApi"
)

type freteClient interface {
	QuoteV3(
		ctx context.Context,
		recipient freteRapidoApi.Recipient,
		dispatchers []freteRapidoApi.DispatcherRequest,
	) (freteRapidoApi.ResponseCotacaoFreteV3, error)
	GetRegisteredNumber() string
}

type Gateway struct {
	client           freteClient
	registeredNumber string
}

func NewFreteGateway(
	client freteClient,
) interfaces.QuoteGateway {
	return &Gateway{
		client:           client,
		registeredNumber: client.GetRegisteredNumber(),
	}
}

func (g *Gateway) Quote(
	ctx context.Context,
	recipientZip int32,
	volumes []entity.VolumeEntity,
) (quotes []entity.QuoteEntity, err error) {

	dispatchers := []freteRapidoApi.DispatcherRequest{
		{
			RegisteredNumber: g.registeredNumber,
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

func (g *Gateway) GetRegisteredNumber() string {
	return g.client.GetRegisteredNumber()
}
