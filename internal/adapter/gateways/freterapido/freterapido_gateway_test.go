package freterapido_test

import (
	"context"
	"errors"
	"reflect"
	"testeFreteRapido/internal/adapter/gateways/freterapido"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/pkg/freteRapidoApi"
	"testing"
)

func TestGateway_Quote(t *testing.T) {
	tests := []struct {
		name         string
		recipientZip int32
		volumes      []entity.VolumeEntity
		mockResp     freteRapidoApi.ResponseCotacaoFreteV3
		mockErr      error
		want         []entity.QuoteEntity
		wantErr      bool
	}{
		{
			name:         "success - one dispatcher one offer",
			recipientZip: 12345678,
			volumes: []entity.VolumeEntity{
				{
					Category:      1,
					Amount:        1,
					UnitaryWeight: 5,
					Price:         100,
					Sku:           "sku123",
					Height:        0.1,
					Width:         0.2,
					Length:        0.3,
					UnitaryPrice:  100,
				},
			},
			mockResp: freteRapidoApi.ResponseCotacaoFreteV3{
				Dispatchers: []freteRapidoApi.DispatcherResponse{
					{
						Offers: []freteRapidoApi.Offer{
							{
								Service:    "SEDEX",
								FinalPrice: 120.50,
								Carrier: freteRapidoApi.Carrier{
									Name: "Correios",
								},
								DeliveryTime: freteRapidoApi.DeliveryTime{Days: 5},
							},
						},
					},
				},
			},
			mockErr: nil,
			want: []entity.QuoteEntity{
				{
					CarrierName:  "Correios",
					Service:      "SEDEX",
					DeadlineDays: 5,
					Price:        120.50,
				},
			},
			wantErr: false,
		},
		{
			name:         "error - API fails",
			recipientZip: 12345678,
			volumes: []entity.VolumeEntity{
				{
					Category:      1,
					Amount:        1,
					UnitaryWeight: 5,
					Price:         100,
					Sku:           "sku123",
					Height:        0.1,
					Width:         0.2,
					Length:        0.3,
					UnitaryPrice:  100,
				},
			},
			mockResp: freteRapidoApi.ResponseCotacaoFreteV3{},
			mockErr:  errors.New("network error"),
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockFreteClient{
				GetRegisteredNumberFunc: func() string { return "1234567890" },
				QuoteV3Func: func(ctx context.Context, recipient freteRapidoApi.Recipient, dispatchers []freteRapidoApi.DispatcherRequest) (freteRapidoApi.ResponseCotacaoFreteV3, error) {
					return tt.mockResp, tt.mockErr
				},
			}

			g := freterapido.NewFreteGateway(mockClient)

			got, err := g.Quote(context.Background(), tt.recipientZip, tt.volumes)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Quote() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Quote() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestGateway_GetRegisteredNumber(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "should return registered number from client",
			want: "1234567890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := freterapido.NewFreteGateway(&freteRapidoApi.Client{
				RegisteredNumber: "1234567890",
			})
			got := g.GetRegisteredNumber()
			if got != tt.want {
				t.Errorf("Gateway.GetRegisteredNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
