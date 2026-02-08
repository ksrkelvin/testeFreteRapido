package quote

import (
	"context"
	"testeFreteRapido/internal/adapter/repositories/mappers"
	"testeFreteRapido/internal/domain/interfaces"
)

type UseCase struct {
	freightGateway interfaces.FreightGateway
	quoteRepo      interfaces.QuoteRepository
}

func NewUseCase(
	freightGateway interfaces.FreightGateway,
	quoteRepo interfaces.QuoteRepository,
) *UseCase {
	return &UseCase{
		freightGateway: freightGateway,
		quoteRepo:      quoteRepo,
	}
}

func (uc *UseCase) Execute(ctx context.Context, input Input) (*Output, error) {
	quotes, err := uc.freightGateway.Quote(ctx, input.RecipientZip, input.Volumes)
	if err != nil {
		return nil, err
	}

	quoteModel := mappers.MapFreightQuotesToModel(quotes)

	err = uc.quoteRepo.SaveQuote(quoteModel)
	if err != nil {
		return nil, err
	}

	return &Output{Quotes: quotes}, nil
}
