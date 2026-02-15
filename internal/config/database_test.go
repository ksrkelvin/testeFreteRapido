package config_test

import (
	"testeFreteRapido/internal/config"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *config.Config
		wantErr bool
	}{
		{
			name:    "erro quando DBConnURL vazio",
			cfg:     &config.Config{DBConnURL: ""},
			wantErr: true,
		},
		{
			name:    "erro quando conexão falha",
			cfg:     &config.Config{DBConnURL: "postgres://invalid:5432/db"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := config.NewDatabase(tt.cfg)
			if tt.wantErr {
				if err == nil {
					t.Fatal("esperado erro, obteve nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}
			if got == nil {
				t.Fatal("esperado gorm.DB, obteve nil")
			}
		})
	}
}
