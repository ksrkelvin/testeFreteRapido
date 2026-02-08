package handlers

import (
	"testeFreteRapido/internal/adapter/http/dtos"
	"testeFreteRapido/internal/adapter/http/mappers"
	"testeFreteRapido/internal/usecase/quote"
	"testeFreteRapido/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type QuoteHandler struct {
	usecase *quote.UseCase
}

func NewQuoteHandler(uc *quote.UseCase) *QuoteHandler {
	return &QuoteHandler{usecase: uc}
}

func (h *QuoteHandler) Quote(c *gin.Context) {
	var req dtos.QuoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		status, body := httpx.FromResult(
			nil,
			httpx.InvalidInput("invalid request body", err.Error()),
		)
		c.JSON(status, body)
		return
	}

	input, err := mappers.ToUseCaseInput(req)
	if err != nil {
		status, body := httpx.FromResult(
			nil,
			httpx.InvalidInput("invalid request data", err.Error()),
		)
		c.JSON(status, body)
		return
	}

	out, err := h.usecase.Execute(c.Request.Context(), input)
	if err != nil {
		status, body := httpx.FromResult(nil, err)
		c.JSON(status, body)
		return
	}

	response := mappers.ToDTOQuotesOutput(out.Quotes)

	status, body := httpx.FromResult(response, nil)
	c.JSON(status, body)
}
