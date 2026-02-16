package main

import (
	"log"

	"testeFreteRapido/internal/adapter/gateways/freterapido"
	"testeFreteRapido/internal/adapter/http"
	"testeFreteRapido/internal/adapter/http/handlers"
	"testeFreteRapido/internal/adapter/repositories"
	"testeFreteRapido/internal/adapter/repositories/migrations"
	"testeFreteRapido/internal/config"
	"testeFreteRapido/internal/usecase/metrics"
	"testeFreteRapido/internal/usecase/quote"
	"testeFreteRapido/pkg/freteRapidoApi"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := cfg.NewDatabase(nil)
	if err != nil {
		log.Fatal(err)
	}

	migrations.Migrate(db)

	r, err := setupServer(cfg, db)
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setupServer(cfg config.AppProvider, db *gorm.DB) (r *gin.Engine, err error) {
	r = gin.Default()

	repo := repositories.NewRepository(db)

	freteClient, err := freteRapidoApi.NewClient(cfg.GetAuthToken(), cfg.GetPlataformCode(), cfg.GetRegisteredNumber())
	if err != nil {
		return nil, err
	}

	freteGateway := freterapido.NewFreteGateway(freteClient)

	quoteUC := quote.NewUseCase(freteGateway, repo.Quote)
	quoteHandler := handlers.NewQuoteHandler(quoteUC)

	metricsUC := metrics.NewUseCase(repo.Metrics)
	metricsHandler := handlers.NewMetricsHandler(metricsUC)

	http.RegisterRoutes(r, quoteHandler, metricsHandler)

	return r, nil
}
