package mappers

import (
	"fmt"
	"strconv"
	"testeFreteRapido/internal/adapter/http/dtos"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/internal/usecase/quote"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ToUseCaseInput(req dtos.QuoteRequestDTO) (input quote.Input, err error) {
	if err := validate.Struct(req); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return quote.Input{}, err
		}

		for _, err := range err.(validator.ValidationErrors) {
			return quote.Input{}, fmt.Errorf("Data '%s' is invalid: %s", err.Namespace(), err.Tag())
		}
	}

	zipCode, err := strconv.Atoi(req.Recipient.Address.Zipcode)
	if err != nil {
		return quote.Input{}, err
	}

	volumes := make([]entity.VolumeEntity, len(req.Volumes))
	for i, v := range req.Volumes {
		volumes[i] = entity.VolumeEntity{
			Category:      v.Category,
			Amount:        v.Amount,
			UnitaryWeight: v.UnitaryWeight,
			Price:         v.Price,
			Sku:           v.Sku,
			Height:        v.Height,
			Width:         v.Width,
			Length:        v.Length,
			UnitaryPrice:  v.UnitaryPrice,
		}
	}

	return quote.Input{
		RecipientZip: int32(zipCode),
		Volumes:      volumes,
	}, nil
}

func ToDTOQuotesOutput(quotes []entity.QuoteEntity) dtos.QuoteResponseDTO {
	var carriers []dtos.CarrierDTO

	for _, q := range quotes {

		carriers = append(carriers, dtos.CarrierDTO{
			Name:     q.CarrierName,
			Service:  q.Service,
			Deadline: strconv.Itoa(int(q.DeadlineDays)),
			Price:    q.Price,
		})
	}

	return dtos.QuoteResponseDTO{
		Carrier: carriers,
	}
}
