package interfaces

import "testeFreteRapido/internal/domain/models"

type QuoteRepository interface {
	SaveQuote(data *models.QuoteModel) (err error)
}
