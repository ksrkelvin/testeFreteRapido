package mappers

import (
	"strconv"
	"testeFreteRapido/internal/adapter/http/dtos"
	"testeFreteRapido/internal/domain/interfaces"
	"testeFreteRapido/internal/usecase/quote"
)

func ToUseCaseInput(req dtos.QuoteRequest) (quote.Input, error) {
	zip, err := strconv.Atoi(req.Recipient.Address.Zipcode)
	if err != nil {
		return quote.Input{}, err
	}

	volumes := make([]interfaces.Volume, len(req.Volumes))
	for i, v := range req.Volumes {
		volumes[i] = interfaces.Volume{
			Category:      v.Category,
			Amount:        v.Amount,
			UnitaryWeight: v.UnitaryWeight,
			Price:         v.Price,
			Sku:           v.Sku,
			Height:        v.Height,
			Width:         v.Width,
			Length:        v.Length,
			UnitaryPrice:  v.Price,
		}
	}

	return quote.Input{
		RecipientZip: int32(zip),
		Volumes:      volumes,
	}, nil
}

func ToDTOOutput(quotes []interfaces.FreightQuote) dtos.QuoteResponse {
	var carriers []dtos.Carrier

	for _, q := range quotes {
		var deadline dtos.Deadline
		deadline.Integer = &q.DeadlineDays
		deadline.String = nil
		carriers = append(carriers, dtos.Carrier{
			Name:     q.CarrierName,
			Service:  q.Service,
			Deadline: &deadline,
			Price:    q.FinalPrice,
		})
	}

	return dtos.QuoteResponse{
		Carrier: carriers,
	}
}
