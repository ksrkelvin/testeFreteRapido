package mappers_test

import (
	"reflect"
	"testing"

	"testeFreteRapido/internal/adapter/http/dtos"
	"testeFreteRapido/internal/adapter/http/mappers"
	"testeFreteRapido/internal/usecase/metrics"
)

func TestToMetricsResponse(t *testing.T) {
	tests := []struct {
		name string
		out  *metrics.Output
		want dtos.MetricsResponseDTO
	}{
		{
			name: "should calculate average correctly when total quotes > 0",
			out: &metrics.Output{
				CarrierMetrics: []metrics.CarrierMetrics{
					{
						CarrierName: "Carrier A",
						TotalQuotes: 2,
						TotalPrice:  200.0,
					},
				},
				CheapestCarrier: metrics.QuoteMetrics{
					CarrierName: "Carrier A",
					Price:       80.0,
				},
				ExpensiveCarrier: metrics.QuoteMetrics{
					CarrierName: "Carrier A",
					Price:       120.0,
				},
			},
			want: dtos.MetricsResponseDTO{
				CarrierMetrics: []dtos.CarrierMetricsDTO{
					{
						CarrierName:  "Carrier A",
						TotalQuotes:  2,
						TotalPrice:   200.0,
						AveragePrice: 100.0,
					},
				},
				CheapestCarrier: dtos.QuoteMetricsDTO{
					CarrierName: "Carrier A",
					Price:       80.0,
				},
				ExpensiveCarrier: dtos.QuoteMetricsDTO{
					CarrierName: "Carrier A",
					Price:       120.0,
				},
			},
		},
		{
			name: "should return zero average when total quotes is zero",
			out: &metrics.Output{
				CarrierMetrics: []metrics.CarrierMetrics{
					{
						CarrierName: "Carrier B",
						TotalQuotes: 0,
						TotalPrice:  0,
					},
				},
				CheapestCarrier:  metrics.QuoteMetrics{},
				ExpensiveCarrier: metrics.QuoteMetrics{},
			},
			want: dtos.MetricsResponseDTO{
				CarrierMetrics: []dtos.CarrierMetricsDTO{
					{
						CarrierName:  "Carrier B",
						TotalQuotes:  0,
						TotalPrice:   0,
						AveragePrice: 0,
					},
				},
				CheapestCarrier:  dtos.QuoteMetricsDTO{},
				ExpensiveCarrier: dtos.QuoteMetricsDTO{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mappers.ToMetricsResponse(tt.out)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMetricsResponse() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
