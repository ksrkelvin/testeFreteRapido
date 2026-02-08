package freterapido

import (
	"strconv"
	"testeFreteRapido/internal/domain/interfaces"
	"testeFreteRapido/pkg/freteRapidoApi"
)

func adaptResponse(resp freteRapidoApi.ResponseCotacaoFreteV3) []interfaces.FreightQuote {
	var quotes []interfaces.FreightQuote

	for _, dispatcher := range resp.Dispatchers {
		for _, offer := range dispatcher.Offers {
			quotes = append(quotes, interfaces.FreightQuote{
				CarrierName:  offer.Carrier.Name,
				Service:      offer.Service,
				DeadlineDays: offer.DeliveryTime.Days,
				FinalPrice:   offer.FinalPrice,
			})
		}
	}

	return quotes
}

func adaptVolumes(vols []interfaces.Volume) []freteRapidoApi.Volume {
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
