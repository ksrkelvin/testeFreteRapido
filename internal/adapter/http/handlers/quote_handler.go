package handlers

import (
	"net/http"
	"testeFreteRapido/internal/adapter/http/dtos"
	"testeFreteRapido/internal/adapter/http/mappers"
	"testeFreteRapido/internal/usecase/quote"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input, err := mappers.ToUseCaseInput(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	out, err := h.usecase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseDTO := mappers.ToDTOOutput(out.Quotes)
	c.JSON(http.StatusOK, responseDTO)
}
