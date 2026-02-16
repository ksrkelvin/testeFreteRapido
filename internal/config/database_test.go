package config_test

import (
	"testeFreteRapido/internal/config"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestNewDatabaseWithMockDriver(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *config.Config
		dbConn  config.DBConn
		wantErr bool
	}{
		{
			name:    "DBConnURL vazio",
			cfg:     &config.Config{DBConnURL: ""},
			dbConn:  mockDriverSuccess,
			wantErr: true,
		},
		{
			name:    "falha na conexão (driver retorna erro)",
			cfg:     &config.Config{DBConnURL: "dummy"},
			dbConn:  mockDriverFail,
			wantErr: true,
		},
		{
			name: "sucesso",
			cfg: &config.Config{
				DBConnURL: "dummy",
				MigrateDB: func(db *gorm.DB) {
				},
			},
			dbConn:  mockDriverSuccess,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			db, err := tt.cfg.NewDatabase(tt.dbConn)
			duration := time.Since(start)

			if tt.wantErr {
				if err == nil {
					t.Fatal("esperado erro, obteve nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}

			if db == nil {
				t.Fatal("esperado gorm.DB, obteve nil")
			}

			t.Logf("teste executado em %v", duration)
		})
	}
}
