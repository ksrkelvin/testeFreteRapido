package handlers

import (
	"context"
	"testeFreteRapido/internal/adapter/http/mappers"
	"testeFreteRapido/internal/usecase/metrics"
	"testeFreteRapido/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type MetricsUsecase interface {
	Execute(ctx context.Context, lastQuotes string) (*metrics.Output, error)
}

type MetricsHandler struct {
	usecase MetricsUsecase
}

func NewMetricsHandler(uc MetricsUsecase) *MetricsHandler {
	return &MetricsHandler{usecase: uc}
}

func (h *MetricsHandler) Metrics(c *gin.Context) {
	param := c.Query("last_quotes")

	out, err := h.usecase.Execute(c.Request.Context(), param)
	if err != nil {
		status, body := httpx.FromResult(nil, err)
		c.JSON(status, body)
		return
	}

	response := mappers.ToMetricsResponse(out)

	status, body := httpx.FromResult(response, nil)
	c.JSON(status, body)
}
