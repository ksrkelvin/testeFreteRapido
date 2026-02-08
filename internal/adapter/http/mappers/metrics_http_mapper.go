package mappers

import (
	"testeFreteRapido/internal/adapter/http/dtos"
	"testeFreteRapido/internal/usecase/metrics"
)

func ToMetricsResponse(out *metrics.Output) dtos.MetricsResponse {
	var carrierMetrics []dtos.CarrierMetrics

	for _, data := range out.CarrierMetrics {
		avg := 0.0
		if data.TotalQuotes > 0 {
			avg = data.TotalPrice / float64(data.TotalQuotes)
		}

		carrierMetrics = append(carrierMetrics, dtos.CarrierMetrics{
			CarrierName:  data.CarrierName,
			TotalQuotes:  data.TotalQuotes,
			TotalPrice:   data.TotalPrice,
			AveragePrice: avg,
		})
	}

	return dtos.MetricsResponse{
		CarrierMetrics: carrierMetrics,
		CheapestCarrier: dtos.QuoteMetrics{
			CarrierName: out.CheapestCarrier.CarrierName,
			Price:       out.CheapestCarrier.Price,
		},
		ExpensiveCarrier: dtos.QuoteMetrics{
			CarrierName: out.ExpensiveCarrier.CarrierName,
			Price:       out.ExpensiveCarrier.Price,
		},
	}
}
