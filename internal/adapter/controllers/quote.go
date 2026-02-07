package controllers

import (
	"testeFreteRapido/internal/usecase/services"

	"github.com/gin-gonic/gin"
)

type cotacaoFreteV3Controller struct {
	cotacaoFreteV3Service services.QuoteService
}

type CotacaoFreteV3Controller interface {
	Router(eng *gin.Engine) (err error)
	PostQuote(ctx *gin.Context)
}

func NewCotacaoFreteV3Controller(cotacaoFreteV3Service services.QuoteService) *cotacaoFreteV3Controller {
	return &cotacaoFreteV3Controller{
		cotacaoFreteV3Service: cotacaoFreteV3Service,
	}
}

func (c *cotacaoFreteV3Controller) Router(eng *gin.Engine) (err error) {
	eng.POST("/quote", c.PostQuote)
	return err
}

func (c *cotacaoFreteV3Controller) PostQuote(ctx *gin.Context) {
	response, err := c.cotacaoFreteV3Service.PostCotacaoFreteV3(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, response)
}
