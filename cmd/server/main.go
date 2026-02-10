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
)

func main() {
	r := gin.Default()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.NewDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	migrations.Migrate(db)

	repo := repositories.NewRepository(db)

	freteClient, err := freteRapidoApi.NewClient(cfg.AuthToken, cfg.PlataformCode, cfg.RegisteredNumber)
	if err != nil {
		log.Fatal(err)
	}

	freteGateway := freterapido.NewFreteGateway(freteClient)

	quoteUC := quote.NewUseCase(freteGateway, repo.Quote)
	quoteHandler := handlers.NewQuoteHandler(quoteUC)

	metricsUC := metrics.NewUseCase(repo.Metrics)
	metricsHandler := handlers.NewMetricsHandler(metricsUC)

	http.RegisterRoutes(r, quoteHandler, metricsHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
