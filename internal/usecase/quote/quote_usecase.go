package quote

import (
	"context"
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

	// Aqui você pode salvar no banco se quiser:
	// uc.quoteRepo.Save(...)

	return &Output{Quotes: quotes}, nil
}
