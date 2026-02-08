package mappers

import (
	"strconv"
	"testeFreteRapido/internal/adapter/repositories/models"
	"testeFreteRapido/internal/domain/entity"
)

func MapFreightQuotesToModel(quotes []entity.QuoteEntity) *models.Quote {
	carriers := make([]models.Carrier, len(quotes))
	for i, q := range quotes {
		carriers[i] = models.Carrier{
			Name:     q.CarrierName,
			Service:  q.Service,
			Deadline: strconv.Itoa(int(q.DeadlineDays)),
			Price:    q.Price,
		}
	}

	return &models.Quote{
		Carriers: carriers,
	}
}
