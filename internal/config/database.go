package config

import (
	"fmt"
	"testeFreteRapido/internal/adapter/repositories/migrations"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (c *Config) NewDatabase(dbCon DBConn) (*gorm.DB, error) {
	if c.DBConnURL == "" {
		return nil, fmt.Errorf("DB_CONN_URL is empty")
	}

	if dbCon == nil {
		dbCon = func(dsn string, cfg *gorm.Config) (*gorm.DB, error) {
			return gorm.Open(postgres.Open(dsn), cfg)
		}
	}

	var db *gorm.DB
	var err error
	const maxRetries = 5
	const retryInterval = 3 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = dbCon(c.DBConnURL, &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
		if err != nil {
			time.Sleep(retryInterval)
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			time.Sleep(retryInterval)
			continue
		}

		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)

		if err := sqlDB.Ping(); err != nil {
			time.Sleep(retryInterval)
			continue
		}

		c.MigrateDatabase(db)

		return db, nil
	}

	return nil, fmt.Errorf("all attempts to connect to database failed: last error: %w", err)
}

func (c *Config) MigrateDatabase(db *gorm.DB) {
	if c.MigrateDB != nil {
		c.MigrateDB(db)
		return
	}
	migrations.Migrate(db)
}
