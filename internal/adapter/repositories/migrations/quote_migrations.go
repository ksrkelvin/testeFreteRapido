package migrations

import (
	"log"
	"testeFreteRapido/internal/domain/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.QuoteModel{},
		&models.CarrierModel{},
	)
	if err != nil {
		log.Fatalf("Falha ao migrar banco: %v", err)
	}
	log.Println("Migrations aplicadas com sucesso")
}
