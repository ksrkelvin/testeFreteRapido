package postgres_test

import (
	"regexp"
	"testeFreteRapido/internal/adapter/repositories/postgres"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/pkg/httpx"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func TestQuoteRepository_SaveQuote(t *testing.T) {
	gormDB, mock := SetupMockDB(t)
	repo := postgres.NewQuoteRepository(gormDB)

	quotes := []entity.QuoteEntity{
		{
			CarrierName: "Carrier A",
			Price:       100.0,
		},
		{
			CarrierName: "Carrier B",
			Price:       150.0,
		},
	}

	tests := []struct {
		name      string
		mockSetup func()
		data      []entity.QuoteEntity
		wantErr   bool
		wantCode  int
	}{
		{
			name: "sucesso ao salvar",
			data: quotes,
			mockSetup: func() {
				mock.ExpectBegin()

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "quotes"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "carriers"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2))

				mock.ExpectCommit()
			},
			wantErr: false,
		},

		{
			name: "erro duplicidade",
			data: quotes,
			mockSetup: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "quotes"`)).
					WillReturnError(gorm.ErrDuplicatedKey)
				mock.ExpectRollback()
			},
			wantErr:  true,
			wantCode: 409,
		},
		{
			name: "erro interno do banco",
			data: quotes,
			mockSetup: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "quotes"`)).
					WillReturnError(gorm.ErrInvalidData)
				mock.ExpectRollback()
			},
			wantErr:  true,
			wantCode: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			err := repo.SaveQuote(tt.data)
			if tt.wantErr {
				if err == nil {
					t.Fatal("esperado erro, obteve nil")
				}
				appErr, ok := err.(*httpx.AppError)
				if !ok {
					t.Fatalf("esperado AppError, obteve %T", err)
				}
				if tt.wantCode != 0 && appErr.StatusCode != tt.wantCode {
					t.Fatalf("esperado status %d, obteve %d", tt.wantCode, appErr.StatusCode)
				}
				return
			}

			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}
		})
	}
}
