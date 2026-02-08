package handlers

import (
	"net/http"
	"testeFreteRapido/internal/usecase/metrics"

	"github.com/gin-gonic/gin"
)

type MetricsHandler struct {
	usecase *metrics.UseCase
}

func NewMetricsHandler(uc *metrics.UseCase) *MetricsHandler {
	return &MetricsHandler{usecase: uc}
}

func (h *MetricsHandler) Metrics(c *gin.Context) {
	err := h.usecase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
