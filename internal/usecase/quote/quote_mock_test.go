package quote

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
}

func (m *MockGateway) Quote(ctx context.Context, zip int32, volumes []entity.VolumeEntity) ([]entity.QuoteEntity, error) {
	return m.QuoteFunc(ctx, zip, volumes)
}

type MockRepository struct {
	SaveQuoteFunc func(quotes []entity.QuoteEntity) error
}

func (m *MockRepository) SaveQuote(quotes []entity.QuoteEntity) error {
	return m.SaveQuoteFunc(quotes)
}
