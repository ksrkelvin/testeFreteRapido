package http

import (
	"testeFreteRapido/internal/adapter/http/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	quoteHandler *handlers.QuoteHandler,
) {
	r.POST("/quote", quoteHandler.Quote)
}
