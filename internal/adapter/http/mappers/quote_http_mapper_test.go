package mappers_test

import (
	"reflect"
	"testeFreteRapido/internal/adapter/http/dtos"
	"testeFreteRapido/internal/adapter/http/mappers"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/internal/usecase/quote"
	"testing"
)

func TestToUseCaseInput(t *testing.T) {
	tests := []struct {
		name    string
		req     dtos.QuoteRequestDTO
		want    quote.Input
		wantErr bool
	}{
		{
			name: "should convert valid request",
			req: dtos.QuoteRequestDTO{
				Recipient: dtos.RecipientDTO{
					Address: dtos.AddressDTO{
						Zipcode: "12345",
					},
				},
				Volumes: []dtos.VolumeDTO{
					{
						Category:      1,
						Amount:        1,
						UnitaryWeight: 2,
						Price:         100,
						Sku:           "SKU1",
						Height:        10,
						Width:         10,
						Length:        10,
						UnitaryPrice:  100,
					},
				},
			},
			want: quote.Input{
				RecipientZip: 12345,
				Volumes: []entity.VolumeEntity{
					{
						Category:      1,
						Amount:        1,
						UnitaryWeight: 2,
						Price:         100,
						Sku:           "SKU1",
						Height:        10,
						Width:         10,
						Length:        10,
						UnitaryPrice:  100,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "should return error when zipcode is invalid",
			req: dtos.QuoteRequestDTO{
				Recipient: dtos.RecipientDTO{
					Address: dtos.AddressDTO{
						Zipcode: "abc",
					},
				},
			},
			wantErr: true,
		},
		{
			name:    "should return validation error when required fields missing",
			req:     dtos.QuoteRequestDTO{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mappers.ToUseCaseInput(tt.req)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestToDTOQuotesOutput(t *testing.T) {
	tests := []struct {
		name   string
		quotes []entity.QuoteEntity
		want   dtos.QuoteResponseDTO
	}{
		{
			name: "should convert single quote",
			quotes: []entity.QuoteEntity{
				{
					CarrierName:  "Carrier A",
					Service:      "Express",
					DeadlineDays: 3,
					Price:        150.0,
				},
			},
			want: dtos.QuoteResponseDTO{
				Carrier: []dtos.CarrierDTO{
					{
						Name:     "Carrier A",
						Service:  "Express",
						Deadline: "3",
						Price:    150.0,
					},
				},
			},
		},
		{
			name:   "should return empty list when no quotes",
			quotes: []entity.QuoteEntity{},
			want: dtos.QuoteResponseDTO{
				Carrier: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mappers.ToDTOQuotesOutput(tt.quotes)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}
