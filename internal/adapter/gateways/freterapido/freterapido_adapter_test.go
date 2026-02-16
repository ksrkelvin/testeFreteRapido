package freterapido

import (
	"errors"
	"reflect"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/pkg/freteRapidoApi"
	"testing"
)

func Test_adaptResponse(t *testing.T) {
	tests := []struct {
		name string
		resp freteRapidoApi.ResponseCotacaoFreteV3
		want []entity.QuoteEntity
	}{
		{
			name: "one dispatcher one offer",
			resp: freteRapidoApi.ResponseCotacaoFreteV3{
				Dispatchers: []freteRapidoApi.DispatcherResponse{
					{
						Offers: []freteRapidoApi.Offer{
							{
								Service:    "SEDEX",
								FinalPrice: 120.50,
								Carrier: freteRapidoApi.Carrier{
									Name: "Correios",
								},
								DeliveryTime: freteRapidoApi.DeliveryTime{
									Days: 5,
								},
							},
						},
					},
				},
			},
			want: []entity.QuoteEntity{
				{
					CarrierName:  "Correios",
					Service:      "SEDEX",
					DeadlineDays: 5,
					Price:        120.50,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := adaptResponse(tt.resp)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("adaptResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adaptVolumes(t *testing.T) {
	tests := []struct {
		name string
		vols []entity.VolumeEntity
		want []freteRapidoApi.Volume
	}{
		{
			name: "single volume",
			vols: []entity.VolumeEntity{
				{
					Category:      7,
					Amount:        2,
					UnitaryWeight: 5,
					Price:         300,
					Sku:           "abc123",
					Height:        0.2,
					Width:         0.3,
					Length:        0.4,
					UnitaryPrice:  150.0,
				},
			},
			want: []freteRapidoApi.Volume{
				{
					Category:      "7",
					Amount:        2,
					UnitaryWeight: 5,
					Price:         300,
					Sku:           "abc123",
					Height:        0.2,
					Width:         0.3,
					Length:        0.4,
					UnitaryPrice:  150.0,
				},
			},
		},
		{
			name: "empty volumes",
			vols: []entity.VolumeEntity{},
			want: []freteRapidoApi.Volume{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := adaptVolumes(tt.vols)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("adaptVolumes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adaptError(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		wantErr bool
	}{
		{
			name:    "generic error",
			err:     errors.New("network error"),
			wantErr: true,
		},
		{
			name: "api 400 error",
			err: &freteRapidoApi.APIError{
				StatusCode: 400,
				Message:    "invalid data",
			},
			wantErr: true,
		},
		{
			name: "api 503 error",
			err: &freteRapidoApi.APIError{
				StatusCode: 503,
				Message:    "service unavailable",
			},
			wantErr: true,
		},
		{
			name: "api 401 error",
			err: &freteRapidoApi.APIError{
				StatusCode: 401,
				Message:    "unauthorized",
			},
			wantErr: true,
		},
		{
			name: "api 403 error",
			err: &freteRapidoApi.APIError{
				StatusCode: 403,
				Message:    "forbidden",
			},
			wantErr: true,
		},
		{
			name: "api 404 error",
			err: &freteRapidoApi.APIError{
				StatusCode: 404,
				Message:    "not found",
			},
			wantErr: true,
		},
		{
			name: "api 429 error",
			err: &freteRapidoApi.APIError{
				StatusCode: 429,
				Message:    "rate limit",
			},
			wantErr: true,
		},
		{
			name: "api 500 error",
			err: &freteRapidoApi.APIError{
				StatusCode: 500,
				Message:    "internal",
			},
			wantErr: true,
		},
		{
			name: "api 502 error",
			err: &freteRapidoApi.APIError{
				StatusCode: 502,
				Message:    "bad gateway",
			},
			wantErr: true,
		},
		{
			name: "api unexpected error",
			err: &freteRapidoApi.APIError{
				StatusCode: 418,
				Message:    "i'm a teapot",
			},
			wantErr: true,
		},
		{
			name:    "nil error",
			err:     nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := adaptError(tt.err)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("adaptError() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("adaptError() succeeded unexpectedly")
			}
		})
	}
}
