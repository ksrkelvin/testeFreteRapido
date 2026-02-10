package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() (config *Config, err error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}
	cfg := &Config{
		DBConnURL:        os.Getenv("DB_CONN_URL"),
		AuthToken:        os.Getenv("FRETE_RAPIDO_API_AUTH_TOKEN"),
		PlataformCode:    os.Getenv("FRETE_RAPIDO_API_PLATAFORM_CODE"),
		RegisteredNumber: REGISTERED_NUMBER,
	}

	if cfg.DBConnURL == "" {
		return nil, fmt.Errorf("DB_CONN_URL is required")
	}

	if cfg.AuthToken == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_AUTH_TOKEN is required")
	}

	if cfg.PlataformCode == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_PLATAFORM_CODE is required")
	}
	return cfg, nil
}
