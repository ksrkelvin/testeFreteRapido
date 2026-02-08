package postgres

import "gorm.io/gorm"

type QuoteRepository struct {
	db *gorm.DB
}

func NewQuoteRepository(db *gorm.DB) *QuoteRepository {
	return &QuoteRepository{
		db: db,
	}
}

func (r *QuoteRepository) SaveQuote(data any) error {
	return nil
}
