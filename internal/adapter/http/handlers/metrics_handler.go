package handlers

import (
	"testeFreteRapido/internal/usecase/metrics"
	"testeFreteRapido/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type MetricsHandler struct {
	usecase *metrics.UseCase
}

func NewMetricsHandler(uc *metrics.UseCase) *MetricsHandler {
	return &MetricsHandler{usecase: uc}
}

func (h *MetricsHandler) Metrics(c *gin.Context) {
	param := c.Query("last_quotes")

	out, err := h.usecase.Execute(c.Request.Context(), param)
	status, body := httpx.FromResult(out, err)
	c.JSON(status, body)
}
