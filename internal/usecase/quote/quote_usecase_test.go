package quote_test

import (
	"context"
	"errors"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/internal/usecase/quote"
	"testing"
)

func TestUseCase_Execute(t *testing.T) {
	tests := []struct {
		name    string
		gateway *MockGateway
		repo    *MockRepository
		input   quote.Input
		want    *quote.Output
		wantErr bool
	}{
		{
			name: "gateway returns error",
			gateway: &MockGateway{
				QuoteFunc: func(ctx context.Context, zip int32, volumes []entity.VolumeEntity) ([]entity.QuoteEntity, error) {
					return nil, errors.New("gateway error")
				},
				GetRegisteredNumberFunc: func() string {
					return "12345678000199"
				},
			},
			repo: &MockRepository{
				SaveQuoteFunc: func(quotes []entity.QuoteEntity) error {
					return nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no quotes returned",
			gateway: &MockGateway{
				QuoteFunc: func(ctx context.Context, zip int32, volumes []entity.VolumeEntity) ([]entity.QuoteEntity, error) {
					return []entity.QuoteEntity{}, nil
				},
				GetRegisteredNumberFunc: func() string {
					return "12345678000199"
				},
			},
			repo: &MockRepository{
				SaveQuoteFunc: func(quotes []entity.QuoteEntity) error {
					return nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "repository returns error",
			gateway: &MockGateway{
				QuoteFunc: func(ctx context.Context, zip int32, volumes []entity.VolumeEntity) ([]entity.QuoteEntity, error) {
					return []entity.QuoteEntity{
						{CarrierName: "Correios", Price: 20},
					}, nil
				},
				GetRegisteredNumberFunc: func() string {
					return "12345678000199"
				},
			},
			repo: &MockRepository{
				SaveQuoteFunc: func(quotes []entity.QuoteEntity) error {
					return errors.New("db error")
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			gateway: &MockGateway{
				QuoteFunc: func(ctx context.Context, zip int32, volumes []entity.VolumeEntity) ([]entity.QuoteEntity, error) {
					return []entity.QuoteEntity{
						{CarrierName: "Correios", Price: 20},
					}, nil
				},
				GetRegisteredNumberFunc: func() string {
					return "12345678000199"
				},
			},
			repo: &MockRepository{
				SaveQuoteFunc: func(quotes []entity.QuoteEntity) error {
					return nil
				},
			},
			want: &quote.Output{
				Quotes: []entity.QuoteEntity{
					{CarrierName: "Correios", Price: 20},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := quote.NewUseCase(tt.gateway, tt.repo)
			got, gotErr := uc.Execute(context.Background(), tt.input)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(got.Quotes) != len(tt.want.Quotes) {
					t.Fatalf("unexpected number of quotes")
				}

				for i := range got.Quotes {
					if got.Quotes[i] != tt.want.Quotes[i] {
						t.Errorf("quote mismatch: got %+v, want %+v",
							got.Quotes[i], tt.want.Quotes[i])
					}
				}
			}
		})
	}
}
