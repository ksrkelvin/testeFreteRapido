package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load() (config *Config, err error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}
	cfg := &Config{
		DBConnURL: os.Getenv("DB_CONN_URL"),

		Host:          os.Getenv("FRETE_RAPIDO_API_HOST"),
		AuthToken:     os.Getenv("FRETE_RAPIDO_API_AUTH_TOKEN"),
		PlataformCode: os.Getenv("FRETE_RAPIDO_API_PLATAFORM_CODE"),
		RegisteredNum: os.Getenv("FRETE_RAPIDO_API_REGISTERED_NUMBER"),
		DispatcherZip: os.Getenv("FRETE_RAPIDO_API_DISPATCHER_ZIP_CODE"),
	}

	if cfg.DBConnURL == "" {
		return nil, fmt.Errorf("DB_CONN_URL is required")
	}

	if cfg.Host == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_HOST is required")
	}

	if cfg.AuthToken == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_AUTH_TOKEN is required")
	}

	if cfg.PlataformCode == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_PLATAFORM_CODE is required")
	}

	if cfg.RegisteredNum == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_REGISTERED_NUMBER is required")
	}

	if cfg.DispatcherZip == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_DISPATCHER_ZIP_CODE is required")
	}

	return cfg, nil
}
