package config

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (dbConfig *Config) ConnectDB() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(
				"Erro to try NewFreteRapidoApi: %v\n\nstack trace:\n%s",
				r,
				debug.Stack(),
			)
		}
	}()
	dbConnUrl := os.Getenv("DB_CONN_URL")

	for i := 0; i < 15; i++ {
		database, err := gorm.Open(postgres.Open(dbConnUrl), &gorm.Config{})
		if err == nil {
			dbConfig.DB = database

			log.Println("Migrating tables")
			if err := database.AutoMigrate(); err != nil {
				log.Fatal("Failed to migrate tables:", err)
			}

			log.Println("Database connection established successfully")
			return nil
		}

		log.Fatalf("Attempt %d: database not available, retrying in 2s... error: %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Could not connect to the database after several attempts")
	return fmt.Errorf("failed to connect to the database")
}
