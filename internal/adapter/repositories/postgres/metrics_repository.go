package postgres

import (
	"errors"
	"strconv"
	"testeFreteRapido/internal/adapter/repositories/mappers"
	"testeFreteRapido/internal/adapter/repositories/models"
	"testeFreteRapido/internal/domain/entity"
	"testeFreteRapido/pkg/httpx"

	"gorm.io/gorm"
)

type MetricsRepository struct {
	db *gorm.DB
}

func NewMetricsRepository(db *gorm.DB) *MetricsRepository {
	return &MetricsRepository{
		db: db,
	}
}

func (r *MetricsRepository) ReadQuotes(lastQuotes string) ([]entity.QuoteEntity, error) {
	var data []models.Quote

	query := r.db.
		Order("created_at DESC").
		Preload("Carriers")

	if lastQuotes != "" {
		limit, err := strconv.Atoi(lastQuotes)
		if err != nil {
			return nil, httpx.InvalidInput(
				"invalid lastQuotes parameter",
				"lastQuotes must be a valid integer",
			)
		}
		query = query.Limit(limit)
	}

	if err := query.Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, httpx.NotFound(
				"quotes not found",
				"no quotes were found",
			)
		}

		return nil, httpx.Internal(err)
	}

	quotes := mappers.MapQuoteModelsToEntities(data)

	return quotes, nil
}
