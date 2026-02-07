package services

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"testeFreteRapido/pkg/freteRapidoApi"
	freteRapidoApiModels "testeFreteRapido/pkg/freteRapidoApi/models"

	"github.com/gin-gonic/gin"
)

type quoteService struct {
	FreteRapidoApi   *freteRapidoApi.FreteRapidoApi
	registeredNumber string
	recipientZipCode int32
}

type QuoteService interface {
	PostCotacaoFreteV3(ctx *gin.Context) (response string, err error)
}

func NewQuoteService(freteRapidoApi *freteRapidoApi.FreteRapidoApi) (service *quoteService, err error) {
	registeredNumber := os.Getenv("FRETE_RAPIDO_API_REGISTERED_NUMBER")
	if registeredNumber == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_REGISTERED_NUMBER environment variable is required")
	}

	recipientZipCodeEnv := os.Getenv("FRETE_RAPIDO_API_RECIPIENT_ZIP_CODE")
	if recipientZipCodeEnv == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_RECIPIENT_ZIP_CODE environment variable is required")
	}
	recipientZipCode, err := strconv.Atoi(recipientZipCodeEnv)
	if err != nil {
		return nil, fmt.Errorf("Error to try parse FRETE_RAPIDO_API_RECIPIENT_ZIP_CODE environment variable: %v", err)
	}

	return &quoteService{
		FreteRapidoApi:   freteRapidoApi,
		registeredNumber: registeredNumber,
		recipientZipCode: int32(recipientZipCode),
	}, nil
}

func (c *quoteService) PostCotacaoFreteV3(ctx *gin.Context) (response string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(
				"Error to try PostCotacaoFreteV3: %v\n\nstack trace:\n%s",
				r,
				debug.Stack(),
			)
		}
	}()

	dispatchers := []freteRapidoApiModels.DispatcherRequest{
		{
			RegisteredNumber: c.registeredNumber,
			Zipcode:          c.recipientZipCode,
			Volumes: []freteRapidoApiModels.Volume{
				{
					Category:      "7",
					Amount:        1,
					UnitaryWeight: 5,
					Price:         349,
					Sku:           "abc-teste-123",
					Height:        0.2,
					Width:         0.2,
					Length:        0.2,
					UnitaryPrice:  349.0,
				},
				{
					Category:      "7",
					Amount:        2,
					UnitaryWeight: 4,
					Price:         556,
					Sku:           "abc-teste-527",
					Height:        0.4,
					Width:         0.6,
					Length:        0.15,
					UnitaryPrice:  278.0,
				},
			},
		},
	}
	recipient := freteRapidoApiModels.Recipient{
		Zipcode: int64(c.recipientZipCode),
	}
	responseCotacaoFreteV3, err := c.FreteRapidoApi.CotacaoFreteV3.PostCotacaoFreteV3(recipient, dispatchers)
	if err != nil {
		return response, err
	}

	responseBytes, err := json.Marshal(responseCotacaoFreteV3)
	if err != nil {
		return response, fmt.Errorf("Error to try marshal response body: %v", err)
	}

	return string(responseBytes), nil
}
