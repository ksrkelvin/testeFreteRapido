package handlers_test

import (
	"context"
	"testeFreteRapido/internal/usecase/metrics"
	"testeFreteRapido/internal/usecase/quote"
)

type MockMetricsUsecase struct {
	ExecuteFunc func(ctx context.Context, lastQuotes string) (*metrics.Output, error)
}

func (m *MockMetricsUsecase) Execute(ctx context.Context, lastQuotes string) (*metrics.Output, error) {
	return m.ExecuteFunc(ctx, lastQuotes)
}

type MockQuoteUsecase struct {
	ExecuteFunc func(ctx context.Context, input quote.Input) (*quote.Output, error)
}

func (m *MockQuoteUsecase) Execute(ctx context.Context, input quote.Input) (*quote.Output, error) {
	return m.ExecuteFunc(ctx, input)
}
