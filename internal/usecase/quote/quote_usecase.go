package quote

import (
	"context"
	"testeFreteRapido/internal/domain/interfaces"
	"testeFreteRapido/pkg/httpx"
)

type UseCase struct {
	gateway interfaces.QuoteGateway
	repo    interfaces.QuoteRepository
}

func NewUseCase(
	gateway interfaces.QuoteGateway,
	repo interfaces.QuoteRepository,
) *UseCase {
	return &UseCase{
		gateway: gateway,
		repo:    repo,
	}
}

func (uc *UseCase) Execute(
	ctx context.Context,
	input Input,
) (*Output, error) {

	quotes, err := uc.gateway.Quote(ctx, input.RecipientZip, input.Volumes)
	if err != nil {
		return nil, httpx.InvalidInput(
			"failed to quote freight",
			err.Error(),
		)
	}

	if len(quotes) == 0 {
		return nil, httpx.NotFound(
			"no quotes found",
			"no carrier returned a valid quote for the given input",
		)
	}

	if err := uc.repo.SaveQuote(quotes); err != nil {
		return nil, httpx.Conflict(
			"failed to save quote",
			err.Error(),
		)
	}

	return &Output{
		Quotes: quotes,
	}, nil
}
