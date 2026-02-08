package repositories

import (
	"testeFreteRapido/internal/adapter/repositories/postgres"
	"testeFreteRapido/internal/domain/interfaces"

	"gorm.io/gorm"
)

type Repository struct {
	Quote   interfaces.QuoteRepository
	Metrics interfaces.MetricRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Quote:   postgres.NewQuoteRepository(db),
		Metrics: postgres.NewMetricsRepository(db),
	}
}
