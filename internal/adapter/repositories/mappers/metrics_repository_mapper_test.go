package mappers_test

import (
	"reflect"
	"testeFreteRapido/internal/adapter/repositories/mappers"
	"testeFreteRapido/internal/adapter/repositories/models"
	"testeFreteRapido/internal/domain/entity"
	"testing"
)

func TestMapQuoteModelsToEntities(t *testing.T) {
	tests := []struct {
		name         string
		modelsQuotes []models.Quote
		want         []entity.QuoteEntity
	}{
		{
			name:         "empty input",
			modelsQuotes: []models.Quote{},
			want:         []entity.QuoteEntity{},
		},
		{
			name: "quote without carriers",
			modelsQuotes: []models.Quote{
				{
					Carriers: []models.Carrier{},
				},
			},
			want: []entity.QuoteEntity{},
		},
		{
			name: "single carrier",
			modelsQuotes: []models.Quote{
				{
					Carriers: []models.Carrier{
						{
							Name:  "Correios",
							Price: 100,
						},
					},
				},
			},
			want: []entity.QuoteEntity{
				{
					CarrierName: "Correios",
					Price:       100,
				},
			},
		},
		{
			name: "multiple carriers in multiple quotes",
			modelsQuotes: []models.Quote{
				{
					Carriers: []models.Carrier{
						{Name: "A", Price: 10},
						{Name: "B", Price: 20},
					},
				},
				{
					Carriers: []models.Carrier{
						{Name: "C", Price: 30},
					},
				},
			},
			want: []entity.QuoteEntity{
				{CarrierName: "A", Price: 10},
				{CarrierName: "B", Price: 20},
				{CarrierName: "C", Price: 30},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mappers.MapQuoteModelsToEntities(tt.modelsQuotes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapQuoteModelsToEntities() = %v, want %v", got, tt.want)
			}
		})
	}
}
