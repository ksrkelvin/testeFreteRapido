package metrics

import "testeFreteRapido/internal/domain/entity"

type MockRepository struct {
	ReadMetricsFunc func(lastQuotes string) (quotes []entity.QuoteEntity, err error)
}

func (m *MockRepository) ReadQuotes(lastQuotes string) (quotes []entity.QuoteEntity, err error) {
	quotes, err = m.ReadMetricsFunc(lastQuotes)
	return quotes, err
}
