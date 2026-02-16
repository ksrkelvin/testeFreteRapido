package interfaces

import "testeFreteRapido/internal/domain/entity"

type QuoteRepository interface {
	SaveQuote(quotes []entity.QuoteEntity) (err error)
}
