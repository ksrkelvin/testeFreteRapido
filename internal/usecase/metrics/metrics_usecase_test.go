package metrics_test

import (
	"context"
	"errors"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/internal/usecase/metrics"
	"testing"
)

func TestUseCase_Execute(t *testing.T) {
	tests := []struct {
		name       string
		repo       *metrics.MockRepository
		lastQuotes string
		want       *metrics.Output
		wantErr    bool
	}{
		{
			name: "repository returns error",
			repo: &metrics.MockRepository{
				ReadMetricsFunc: func(lastQuotes string) ([]entity.QuoteEntity, error) {
					return nil, errors.New("db error")
				},
			},
			wantErr: true,
		},
		{
			name: "no quotes found",
			repo: &metrics.MockRepository{
				ReadMetricsFunc: func(lastQuotes string) ([]entity.QuoteEntity, error) {
					return []entity.QuoteEntity{}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "success",
			repo: &metrics.MockRepository{
				ReadMetricsFunc: func(lastQuotes string) ([]entity.QuoteEntity, error) {
					return []entity.QuoteEntity{
						{CarrierName: "Correios", Price: 10},
						{CarrierName: "Correios", Price: 20},
						{CarrierName: "FedEx", Price: 50},
					}, nil
				},
			},
			want: &metrics.Output{
				CarrierMetrics: []metrics.CarrierMetrics{
					{
						CarrierName:  "Correios",
						TotalQuotes:  2,
						TotalPrice:   30,
						AveragePrice: 15,
					},
					{
						CarrierName:  "FedEx",
						TotalQuotes:  1,
						TotalPrice:   50,
						AveragePrice: 50,
					},
				},
				CheapestCarrier: metrics.QuoteMetrics{
					CarrierName: "Correios",
					Price:       10,
				},
				ExpensiveCarrier: metrics.QuoteMetrics{
					CarrierName: "FedEx",
					Price:       50,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := metrics.NewUseCase(tt.repo)
			got, err := uc.Execute(context.Background(), tt.lastQuotes)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}
			if len(got.CarrierMetrics) != len(tt.want.CarrierMetrics) {
				t.Fatalf("unexpected number of carrier metrics: got %d, want %d",
					len(got.CarrierMetrics), len(tt.want.CarrierMetrics))
			}
			for _, wantCarrier := range tt.want.CarrierMetrics {
				var found bool

				for _, gotCarrier := range got.CarrierMetrics {
					if gotCarrier.CarrierName == wantCarrier.CarrierName {
						found = true

						if gotCarrier.TotalQuotes != wantCarrier.TotalQuotes {
							t.Errorf("TotalQuotes mismatch for %s", wantCarrier.CarrierName)
						}

						if gotCarrier.TotalPrice != wantCarrier.TotalPrice {
							t.Errorf("TotalPrice mismatch for %s", wantCarrier.CarrierName)
						}

						if gotCarrier.AveragePrice != wantCarrier.AveragePrice {
							t.Errorf("AveragePrice mismatch for %s", wantCarrier.CarrierName)
						}
					}
				}

				if !found {
					t.Errorf("carrier %s not found", wantCarrier.CarrierName)
				}
			}
		})
	}
}
