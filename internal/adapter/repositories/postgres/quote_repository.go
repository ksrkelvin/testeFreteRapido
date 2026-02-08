package postgres

import (
	"testeFreteRapido/internal/domain/models"

	"gorm.io/gorm"
)

type QuoteRepository struct {
	db *gorm.DB
}

func NewQuoteRepository(db *gorm.DB) *QuoteRepository {
	return &QuoteRepository{
		db: db,
	}
}

func (r *QuoteRepository) SaveQuote(data *models.QuoteModel) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(data).Error; err != nil {
			return err
		}
		return nil
	})
}
