package mappers

import (
	"testeFreteRapido/internal/adapter/http/dtos"
	"testeFreteRapido/internal/usecase/metrics"
)

func ToMetricsResponse(out *metrics.Output) dtos.MetricsResponseDTO {
	var carrierMetrics []dtos.CarrierMetricsDTO

	for _, data := range out.CarrierMetrics {
		avg := 0.0
		if data.TotalQuotes > 0 {
			avg = data.TotalPrice / float64(data.TotalQuotes)
		}

		carrierMetrics = append(carrierMetrics, dtos.CarrierMetricsDTO{
			CarrierName:  data.CarrierName,
			TotalQuotes:  data.TotalQuotes,
			TotalPrice:   data.TotalPrice,
			AveragePrice: avg,
		})
	}

	return dtos.MetricsResponseDTO{
		CarrierMetrics: carrierMetrics,
		CheapestCarrier: dtos.QuoteMetricsDTO{
			CarrierName: out.CheapestCarrier.CarrierName,
			Price:       out.CheapestCarrier.Price,
		},
		ExpensiveCarrier: dtos.QuoteMetricsDTO{
			CarrierName: out.ExpensiveCarrier.CarrierName,
			Price:       out.ExpensiveCarrier.Price,
		},
	}
}
