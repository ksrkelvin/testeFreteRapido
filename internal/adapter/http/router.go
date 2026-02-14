package http

import (
	"testeFreteRapido/internal/adapter/http/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	quoteHandler *handlers.QuoteHandler,
	metricsHandler *handlers.MetricsHandler,
) {
	r.POST("/quote", quoteHandler.Quote)
	r.GET("/metrics", metricsHandler.Metrics)
}
