package config_test

import (
	"os"
	"testeFreteRapido/internal/config"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name          string
		env           map[string]string
		wantErr       bool
		expectedError string
	}{
		{
			name:          "sucesso com todas variáveis definidas",
			env:           map[string]string{"DB_CONN_URL": "postgres://user:pass@localhost/db", "FRETE_RAPIDO_API_AUTH_TOKEN": "token", "FRETE_RAPIDO_API_PLATAFORM_CODE": "123", "REGISTERED_NUMBER": "25438296000158"},
			wantErr:       false,
			expectedError: "",
		},
		{
			name:          "erro quando DB_CONN_URL ausente",
			env:           map[string]string{"FRETE_RAPIDO_API_AUTH_TOKEN": "token", "FRETE_RAPIDO_API_PLATAFORM_CODE": "123", "REGISTERED_NUMBER": "25438296000158"},
			wantErr:       true,
			expectedError: "DB_CONN_URL is required",
		},
		{
			name:          "erro quando AuthToken ausente",
			env:           map[string]string{"DB_CONN_URL": "postgres://user:pass@localhost/db", "FRETE_RAPIDO_API_PLATAFORM_CODE": "123", "REGISTERED_NUMBER": "25438296000158"},
			wantErr:       true,
			expectedError: "FRETE_RAPIDO_API_AUTH_TOKEN is required",
		},
		{
			name:          "erro quando PlataformCode ausente",
			env:           map[string]string{"DB_CONN_URL": "postgres://user:pass@localhost/db", "FRETE_RAPIDO_API_AUTH_TOKEN": "token", "REGISTERED_NUMBER": "25438296000158"},
			wantErr:       true,
			expectedError: "FRETE_RAPIDO_API_PLATAFORM_CODE is required",
		},
		{
			name:          "erro quando RegisteredNumber ausente",
			env:           map[string]string{"DB_CONN_URL": "postgres://user:pass@localhost/db", "FRETE_RAPIDO_API_AUTH_TOKEN": "token", "FRETE_RAPIDO_API_PLATAFORM_CODE": "123"},
			wantErr:       true,
			expectedError: "REGISTERED_NUMBER is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("DB_CONN_URL")
			os.Unsetenv("FRETE_RAPIDO_API_AUTH_TOKEN")
			os.Unsetenv("FRETE_RAPIDO_API_PLATAFORM_CODE")
			os.Unsetenv("REGISTERED_NUMBER")
			for k, v := range tt.env {
				os.Setenv(k, v)
			}

			got, err := config.Load()
			if tt.wantErr {
				if err == nil {
					t.Fatal("esperado erro, obteve nil")
				}
				if err.Error() != tt.expectedError {
					t.Fatalf("esperado erro %q, obteve %q", tt.expectedError, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("erro inesperado: %v", err)
			}
			if got.GetDBConnURL() != tt.env["DB_CONN_URL"] ||
				got.GetAuthToken() != tt.env["FRETE_RAPIDO_API_AUTH_TOKEN"] ||
				got.GetPlataformCode() != tt.env["FRETE_RAPIDO_API_PLATAFORM_CODE"] ||
				got.GetRegisteredNumber() != tt.env["REGISTERED_NUMBER"] {
				t.Errorf("config carregada incorreta: %+v", got)
			}
		})
	}
}
