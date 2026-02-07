package controllers

import (
	"fmt"
	"runtime/debug"
	"testeFreteRapido/internal/usecase/services"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	CotacaoFreteV3Controller *cotacaoFreteV3Controller
}

func RegisterControllers(eng *gin.Engine, service *services.Service) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(
				"Erro to try NewFreteRapidoApi: %v\n\nstack trace:\n%s",
				r,
				debug.Stack(),
			)
		}
	}()

	c := &Controllers{
		CotacaoFreteV3Controller: NewCotacaoFreteV3Controller(service.QuoteService),
	}

	if err = c.CotacaoFreteV3Controller.Router(eng); err != nil {
		return fmt.Errorf("Error to set routes: %v", err)
	}

	return err
}
