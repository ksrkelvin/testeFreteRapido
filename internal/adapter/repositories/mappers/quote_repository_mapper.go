package mappers

import (
	"strconv"
	"testeFreteRapido/internal/domain/interfaces"
	"testeFreteRapido/internal/domain/models"
)

func MapFreightQuotesToModel(quotes []interfaces.FreightQuote) *models.QuoteModel {
	carriers := make([]models.CarrierModel, len(quotes))
	for i, q := range quotes {
		carriers[i] = models.CarrierModel{
			Name:     q.CarrierName,
			Service:  q.Service,
			Deadline: strconv.Itoa(int(q.DeadlineDays)),
			Price:    q.FinalPrice,
		}
	}

	return &models.QuoteModel{
		Carrier: carriers,
	}
}
