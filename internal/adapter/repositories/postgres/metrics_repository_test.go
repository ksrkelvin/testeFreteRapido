package postgres_test

import (
	"testeFreteRapido/internal/adapter/repositories/postgres"
	"testeFreteRapido/pkg/httpx"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func TestMetricsRepository_ReadQuotes(t *testing.T) {
	gormDB, mock := SetupMockDB(t)
	repo := postgres.NewMetricsRepository(gormDB)

	tests := []struct {
		name       string
		lastQuotes string
		mockSetup  func()
		wantLen    int
		wantErr    bool
		wantCode   int
	}{
		{
			name:       "sucesso sem lastQuotes",
			lastQuotes: "",
			mockSetup: func() {
				quotesRows := sqlmock.NewRows([]string{"id", "created_at"}).
					AddRow(1, time.Date(2026, 2, 15, 10, 0, 0, 0, time.UTC)).
					AddRow(2, time.Date(2026, 2, 15, 9, 0, 0, 0, time.UTC))

				mock.ExpectQuery(`SELECT \* FROM "quotes".*ORDER BY created_at DESC`).
					WillReturnRows(quotesRows)

				carriersRows := sqlmock.NewRows([]string{"id", "quote_id", "name"}).
					AddRow(1, 1, "Carrier A").
					AddRow(2, 2, "Carrier B")

				mock.ExpectQuery(`SELECT \* FROM "carriers".*WHERE "carriers"."quote_id".*`).
					WillReturnRows(carriersRows)
			},
			wantLen: 2,
			wantErr: false,
		},
		{
			name:       "sucesso com lastQuotes válido",
			lastQuotes: "1",
			mockSetup: func() {
				quotesRows := sqlmock.NewRows([]string{"id", "created_at"}).
					AddRow(1, time.Date(2026, 2, 15, 10, 0, 0, 0, time.UTC))

				mock.ExpectQuery(`SELECT \* FROM "quotes".*ORDER BY created_at DESC`).
					WithArgs(1).
					WillReturnRows(quotesRows)

				carriersRows := sqlmock.NewRows([]string{"id", "quote_id", "name"}).
					AddRow(1, 1, "Carrier A")

				mock.ExpectQuery(`SELECT \* FROM "carriers".*WHERE "carriers"."quote_id".*`).
					WithArgs(1).
					WillReturnRows(carriersRows)
			},
			wantLen: 1,
			wantErr: false,
		},
		{
			name:       "erro lastQuotes inválido",
			lastQuotes: "abc",
			mockSetup:  func() {},
			wantErr:    true,
			wantCode:   400,
		},
		{
			name:       "erro do banco",
			lastQuotes: "",
			mockSetup: func() {
				mock.ExpectQuery(`SELECT \* FROM "quotes".*ORDER BY created_at DESC`).
					WillReturnError(gorm.ErrRecordNotFound)
			},
			wantErr:  true,
			wantCode: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			got, err := repo.ReadQuotes(tt.lastQuotes)
			if tt.wantErr {
				if err == nil {
					t.Fatal("esperado erro, obteve nil")
				}
				httpErr, ok := err.(*httpx.AppError)
				if !ok {
					t.Fatalf("esperado AppError, obteve %T", err)
				}
				if httpErr.StatusCode != tt.wantCode {
					t.Fatalf("esperado status %d, obteve %d", tt.wantCode, httpErr.StatusCode)
				}
				return
			}

			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}

			if len(got) != tt.wantLen {
				t.Errorf("esperado %d quotes, obteve %d", tt.wantLen, len(got))
			}
		})
	}
}
