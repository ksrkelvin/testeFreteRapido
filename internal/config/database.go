package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(cfg *Config) (*gorm.DB, error) {
	if cfg.DBConnURL == "" {
		return nil, fmt.Errorf("DB_CONN_URL is empty")
	}

	var db *gorm.DB
	var err error

	const maxRetries = 5
	const retryInterval = 3 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(cfg.DBConnURL), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			fmt.Printf("Attempt %d/%d: failed to connect database: %v\n", i+1, maxRetries, err)
			time.Sleep(retryInterval)
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			fmt.Printf("Attempt %d/%d: failed to get sql.DB from gorm: %v\n", i+1, maxRetries, err)
			time.Sleep(retryInterval)
			continue
		}

		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)

		if err := sqlDB.Ping(); err != nil {
			fmt.Printf("Attempt %d/%d: database ping failed: %v\n", i+1, maxRetries, err)
			time.Sleep(retryInterval)
			continue
		}

		return db, nil
	}

	return nil, fmt.Errorf("all %d attempts to connect to database failed: last error: %w", maxRetries, err)
}
