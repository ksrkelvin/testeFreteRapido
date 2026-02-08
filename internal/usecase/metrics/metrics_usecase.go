package metrics

import (
	"context"
	"testeFreteRapido/internal/domain/interfaces"
)

type UseCase struct {
	quoteRepo interfaces.QuoteRepository
}

func NewUseCase(
	quoteRepo interfaces.QuoteRepository,
) *UseCase {
	return &UseCase{
		quoteRepo: quoteRepo,
	}
}

func (uc *UseCase) Execute(ctx context.Context) (err error) {
	return nil
}
