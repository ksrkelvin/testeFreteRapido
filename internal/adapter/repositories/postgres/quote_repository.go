package postgres

import (
	"errors"
	"testeFreteRapido/internal/adapter/repositories/mappers"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/pkg/httpx"

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

func (r *QuoteRepository) SaveQuote(data []entity.QuoteEntity) error {
	quoteModel := mappers.MapFreightQuotesToModel(data)

	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(quoteModel).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return httpx.Conflict(
					"quote already exists",
					"a quote with the same identifier already exists",
				)
			}

			return err
		}
		return nil
	}); err != nil {

		if _, ok := err.(*httpx.AppError); ok {
			return err
		}

		return httpx.Internal(err)
	}

	return nil
}
