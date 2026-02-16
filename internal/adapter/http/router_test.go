package http_test

import (
	"testeFreteRapido/internal/adapter/http"
	"testeFreteRapido/internal/adapter/http/handlers"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterRoutes(t *testing.T) {
	tests := []struct {
		name           string
		r              *gin.Engine
		quoteHandler   *handlers.QuoteHandler
		metricsHandler *handlers.MetricsHandler
	}{
		{
			name:           "basic route registration",
			r:              gin.Default(),
			quoteHandler:   &handlers.QuoteHandler{},
			metricsHandler: &handlers.MetricsHandler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			http.RegisterRoutes(tt.r, tt.quoteHandler, tt.metricsHandler)
		})
	}
}
