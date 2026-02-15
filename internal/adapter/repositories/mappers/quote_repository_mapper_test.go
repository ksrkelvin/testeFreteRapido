package mappers_test

import (
	"reflect"
	"testeFreteRapido/internal/adapter/repositories/mappers"
	"testeFreteRapido/internal/adapter/repositories/models"
	"testeFreteRapido/internal/domain/entity"
	"testing"
)

func TestMapFreightQuotesToModel(t *testing.T) {
	tests := []struct {
		name   string
		quotes []entity.QuoteEntity
		want   *models.Quote
	}{
		{
			name:   "empty input",
			quotes: []entity.QuoteEntity{},
			want: &models.Quote{
				Carriers: []models.Carrier{},
			},
		},
		{
			name: "single quote",
			quotes: []entity.QuoteEntity{
				{
					CarrierName:  "Correios",
					Service:      "SEDEX",
					DeadlineDays: 3,
					Price:        100.50,
				},
			},
			want: &models.Quote{
				Carriers: []models.Carrier{
					{
						Name:     "Correios",
						Service:  "SEDEX",
						Deadline: "3",
						Price:    100.50,
					},
				},
			},
		},
		{
			name: "multiple quotes preserve order",
			quotes: []entity.QuoteEntity{
				{
					CarrierName:  "A",
					Service:      "S1",
					DeadlineDays: 1,
					Price:        10,
				},
				{
					CarrierName:  "B",
					Service:      "S2",
					DeadlineDays: 5,
					Price:        20,
				},
			},
			want: &models.Quote{
				Carriers: []models.Carrier{
					{
						Name:     "A",
						Service:  "S1",
						Deadline: "1",
						Price:    10,
					},
					{
						Name:     "B",
						Service:  "S2",
						Deadline: "5",
						Price:    20,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mappers.MapFreightQuotesToModel(tt.quotes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapFreightQuotesToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
