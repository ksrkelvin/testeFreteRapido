package main

import (
	"log"

	"testeFreteRapido/internal/adapter/gateways/freterapido"
	"testeFreteRapido/internal/adapter/http"
	"testeFreteRapido/internal/adapter/http/handlers"
	"testeFreteRapido/internal/adapter/repositories"
	"testeFreteRapido/internal/adapter/repositories/migrations"
	"testeFreteRapido/internal/config"
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

	freteClient, err := freteRapidoApi.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	freteGateway := freterapido.NewFreteGateway(freteClient)

	quoteUC := quote.NewUseCase(freteGateway, repo.Quote)

	quoteHandler := handlers.NewQuoteHandler(quoteUC)

	http.RegisterRoutes(r, quoteHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
