package freterapido

import (
	"errors"
	"strconv"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/pkg/freteRapidoApi"
	"testeFreteRapido/pkg/httpx"
)

func adaptResponse(resp freteRapidoApi.ResponseCotacaoFreteV3) []entity.QuoteEntity {
	var quotes []entity.QuoteEntity

	for _, dispatcher := range resp.Dispatchers {
		for _, offer := range dispatcher.Offers {
			quotes = append(quotes, entity.QuoteEntity{
				CarrierName:  offer.Carrier.Name,
				Service:      offer.Service,
				DeadlineDays: int(offer.DeliveryTime.Days),
				Price:        offer.FinalPrice,
			})
		}
	}

	return quotes
}

func adaptVolumes(vols []entity.VolumeEntity) []freteRapidoApi.Volume {
	result := make([]freteRapidoApi.Volume, len(vols))
	for i, v := range vols {
		result[i] = freteRapidoApi.Volume{
			Category:      strconv.Itoa(int(v.Category)),
			Amount:        int64(v.Amount),
			UnitaryWeight: int64(v.UnitaryWeight),
			Price:         int64(v.Price),
			Sku:           v.Sku,
			Height:        v.Height,
			Width:         v.Width,
			Length:        v.Length,
			UnitaryPrice:  v.UnitaryPrice,
		}
	}
	return result
}

func adaptError(err error) error {
	if err == nil {
		return nil
	}
	var apiErr *freteRapidoApi.APIError

	if !errors.As(err, &apiErr) {
		return httpx.BadGateway(
			"frete rapido integration error",
			err.Error(),
		)
	}

	switch apiErr.StatusCode {

	case 400, 422:
		return httpx.BadRequest(
			"invalid request to frete rapido",
			apiErr.Message,
		)

	case 401, 403:
		return httpx.BadGateway(
			"frete rapido authorization error",
			apiErr.Message,
		)

	case 404:
		return httpx.BadGateway(
			"frete rapido resource not found",
			apiErr.Message,
		)

	case 429:
		return httpx.ServiceUnavailable(
			"frete rapido rate limit exceeded",
			apiErr.Message,
		)

	case 500, 502:
		return httpx.BadGateway(
			"frete rapido internal error",
			apiErr.Message,
		)

	case 503:
		return httpx.ServiceUnavailable(
			"frete rapido unavailable",
			apiErr.Message,
		)

	default:
		return httpx.BadGateway(
			"frete rapido unexpected error",
			apiErr.Message,
		)
	}
}
