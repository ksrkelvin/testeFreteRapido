package services

import (
	"testeFreteRapido/internal/adapter/repositories"
	"testeFreteRapido/pkg/freteRapidoApi"
)

type Service struct {
	QuoteService QuoteService
}

func NewService(repo *repositories.Repository) (service *Service, err error) {
	freteRapidoApi, err := freteRapidoApi.NewFreteRapidoApi()
	if err != nil {
		return nil, err
	}

	quoteService, err := NewQuoteService(freteRapidoApi)
	if err != nil {
		return nil, err
	}

	return &Service{
		QuoteService: quoteService,
	}, nil
}
