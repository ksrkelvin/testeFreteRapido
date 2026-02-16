package freterapido_test

import (
	"context"
	"testeFreteRapido/pkg/freteRapidoApi"
)

type MockFreteClient struct {
	QuoteV3Func             func(ctx context.Context, recipient freteRapidoApi.Recipient, dispatchers []freteRapidoApi.DispatcherRequest) (freteRapidoApi.ResponseCotacaoFreteV3, error)
	GetRegisteredNumberFunc func() string
}

func (m *MockFreteClient) QuoteV3(
	ctx context.Context,
	recipient freteRapidoApi.Recipient,
	dispatchers []freteRapidoApi.DispatcherRequest,
) (freteRapidoApi.ResponseCotacaoFreteV3, error) {
	return m.QuoteV3Func(ctx, recipient, dispatchers)
}

func (m *MockFreteClient) GetRegisteredNumber() string {
	return m.GetRegisteredNumberFunc()
}
