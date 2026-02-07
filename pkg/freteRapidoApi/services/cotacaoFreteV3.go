package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"testeFreteRapido/pkg/freteRapidoApi/models"
)

type cotacaoFreteV3 struct {
	utils            *utils
	registeredNumber string
	authToken        string
	plataformCode    string
}

func NewCotacaoFreteV3(registeredNumber string, authToken string, plataformCode string) (cotacaoFrete *cotacaoFreteV3, err error) {
	host := os.Getenv("FRETE_RAPIDO_API_HOST")
	if host == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_HOST environment variable is required")
	}

	utils := NewUtils(host)
	return &cotacaoFreteV3{
		utils:            utils,
		registeredNumber: registeredNumber,
		authToken:        authToken,
		plataformCode:    plataformCode,
	}, nil
}

type CotacaoFreteV3 interface {
	PostCotacaoFreteV3(recipient models.Recipient, dispatchers []models.DispatcherRequest) (response models.ResponseCotacaoFreteV3, err error)
}

func (c *cotacaoFreteV3) PostCotacaoFreteV3(recipient models.Recipient, dispatchers []models.DispatcherRequest) (response models.ResponseCotacaoFreteV3, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(
				"Error to try PostCotacaoFreteV3: %v\n\nstack trace:\n%s",
				r,
				debug.Stack(),
			)
		}
	}()

	reqBody := models.RequestCotacaoFreteV3{
		Shipper: models.Shipper{
			RegisteredNumber: c.registeredNumber,
			Token:            c.authToken,
			PlatformCode:     c.plataformCode,
		},
		Recipient:      recipient,
		Dispatchers:    dispatchers,
		SimulationType: []int{0},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	path := "/api/v3/quote/simulate"

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return response, fmt.Errorf("Error to try marshal request body: %v", err)
	}

	res, statusHttp, err := c.utils.ReqHTTP(path, http.MethodPost, headers, payload)
	if err != nil {
		return response, fmt.Errorf("Error to try make request: %v", err)
	}

	if statusHttp != http.StatusOK {
		return response, fmt.Errorf("Error response from API: %s, http status: %d", res, statusHttp)
	}

	responseBody := models.ResponseCotacaoFreteV3{}
	err = json.Unmarshal(res, &responseBody)
	if err != nil {
		return response, fmt.Errorf("Error to try unmarshal response body: %v", err)
	}

	return responseBody, nil
}
