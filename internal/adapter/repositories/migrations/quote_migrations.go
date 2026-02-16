package migrations

import (
	"log"
	"testeFreteRapido/internal/adapter/repositories/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Quote{},
		&models.Carrier{},
	)
	if err != nil {
		log.Fatalf("Falha ao migrar banco: %v", err)
	}
	log.Println("Migrations aplicadas com sucesso")
}
