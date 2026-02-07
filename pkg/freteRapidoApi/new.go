package freteRapidoApi

import (
	"fmt"
	"os"
	"runtime/debug"
	usecase "testeFreteRapido/pkg/freteRapidoApi/services"
)

type FreteRapidoApi struct {
	CotacaoFreteV3 usecase.CotacaoFreteV3
}

func NewFreteRapidoApi() (freteRapidoApi *FreteRapidoApi, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(
				"Erro to try NewFreteRapidoApi: %v\n\nstack trace:\n%s",
				r,
				debug.Stack(),
			)
		}
	}()

	registeredNumber := os.Getenv("FRETE_RAPIDO_API_REGISTERED_NUMBER")
	if registeredNumber == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_REGISTERED_NUMBER environment variable is required")
	}

	authToken := os.Getenv("FRETE_RAPIDO_API_AUTH_TOKEN")
	if authToken == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_AUTH_TOKEN environment variable is required")
	}

	plataformCode := os.Getenv("FRETE_RAPIDO_API_PLATAFORM_CODE")
	if plataformCode == "" {
		return nil, fmt.Errorf("FRETE_RAPIDO_API_PLATAFORM_CODE environment variable is required")
	}

	cotacaFreteV3, err := usecase.NewCotacaoFreteV3(registeredNumber, authToken, plataformCode)
	if err != nil {
		return nil, fmt.Errorf("Error to try create CotacaoFreteV3: %v", err)
	}

	return &FreteRapidoApi{
		CotacaoFreteV3: cotacaFreteV3,
	}, nil
}
