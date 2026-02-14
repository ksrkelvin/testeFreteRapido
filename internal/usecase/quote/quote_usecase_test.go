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
		gateway *quote.MockGateway
		repo    *quote.MockRepository
		input   quote.Input
		want    *quote.Output
		wantErr bool
	}{
		{
			name: "gateway returns error",
			gateway: &quote.MockGateway{
				QuoteFunc: func(ctx context.Context, zip int32, volumes []entity.VolumeEntity) ([]entity.QuoteEntity, error) {
					return nil, errors.New("gateway error")
				},
			},
			repo: &quote.MockRepository{
				SaveQuoteFunc: func(quotes []entity.QuoteEntity) error {
					return nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no quotes returned",
			gateway: &quote.MockGateway{
				QuoteFunc: func(ctx context.Context, zip int32, volumes []entity.VolumeEntity) ([]entity.QuoteEntity, error) {
					return []entity.QuoteEntity{}, nil
				},
			},
			repo: &quote.MockRepository{
				SaveQuoteFunc: func(quotes []entity.QuoteEntity) error {
					return nil
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "repository returns error",
			gateway: &quote.MockGateway{
				QuoteFunc: func(ctx context.Context, zip int32, volumes []entity.VolumeEntity) ([]entity.QuoteEntity, error) {
					return []entity.QuoteEntity{
						{CarrierName: "Correios", Price: 20},
					}, nil
				},
			},
			repo: &quote.MockRepository{
				SaveQuoteFunc: func(quotes []entity.QuoteEntity) error {
					return errors.New("db error")
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			gateway: &quote.MockGateway{
				QuoteFunc: func(ctx context.Context, zip int32, volumes []entity.VolumeEntity) ([]entity.QuoteEntity, error) {
					return []entity.QuoteEntity{
						{CarrierName: "Correios", Price: 20},
					}, nil
				},
			},
			repo: &quote.MockRepository{
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
