package mappers

import (
	"testeFreteRapido/internal/adapter/repositories/models"
	"testeFreteRapido/internal/domain/entity"
)

func MapQuoteModelsToEntities(modelsQuotes []models.Quote) []entity.QuoteEntity {
	var entities []entity.QuoteEntity

	for _, quote := range modelsQuotes {
		for _, carrier := range quote.Carriers {
			entities = append(entities, entity.QuoteEntity{
				CarrierName: carrier.Name,
				Price:       carrier.Price,
			})
		}
	}

	return entities
}
