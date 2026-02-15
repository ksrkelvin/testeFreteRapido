package quote_test

import (
	"context"
	"testeFreteRapido/internal/domain/entity"
)

type MockGateway struct {
	QuoteFunc func(
		ctx context.Context,
		recipientZip int32,
		volumes []entity.VolumeEntity,
	) (quoteEntity []entity.QuoteEntity, err error)
	GetRegisteredNumberFunc func() string
}

func (m *MockGateway) Quote(ctx context.Context, zip int32, volumes []entity.VolumeEntity) ([]entity.QuoteEntity, error) {
	return m.QuoteFunc(ctx, zip, volumes)
}

func (m *MockGateway) GetRegisteredNumber() string {
	return m.GetRegisteredNumberFunc()
}

type MockRepository struct {
	SaveQuoteFunc func(quotes []entity.QuoteEntity) error
}

func (m *MockRepository) SaveQuote(quotes []entity.QuoteEntity) error {
	return m.SaveQuoteFunc(quotes)
}
