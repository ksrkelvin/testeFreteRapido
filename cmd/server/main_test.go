package main

import (
	"testing"

	"testeFreteRapido/internal/config"

	"gorm.io/gorm"
)

func TestSetupServerBasic(t *testing.T) {
	tests := []struct {
		name string
		cfg  *config.Config
		db   *gorm.DB
	}{
		{
			name: "basic server setup",
			cfg: &config.Config{
				DBConnURL:        "dummy-db",
				AuthToken:        "dummy-token",
				PlataformCode:    "dummy-code",
				RegisteredNumber: "123",
			},
			db: &gorm.DB{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine, err := setupServer(tt.cfg, tt.db)
			if err != nil {
				t.Errorf("setupServer() failed: %v", err)
			}

			if engine == nil {
				t.Error("Expected engine not to be nil")
			}

			if len(engine.Routes()) == 0 {
				t.Error("Expected at least one route to be registered")
			}
		})
	}
}
