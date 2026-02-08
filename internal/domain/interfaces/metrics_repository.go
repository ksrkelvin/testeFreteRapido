package interfaces

import "testeFreteRapido/internal/domain/entity"

type MetricRepository interface {
	ReadQuotes(lastQuotes string) (quotes []entity.QuoteEntity, err error)
}
