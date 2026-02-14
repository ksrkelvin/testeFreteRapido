package metrics

import (
	"context"
	"math"

	"testeFreteRapido/internal/domain/interfaces"
	"testeFreteRapido/pkg/httpx"
)

type UseCase struct {
	repo interfaces.MetricRepository
}

func NewUseCase(repo interfaces.MetricRepository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (uc *UseCase) Execute(
	ctx context.Context,
	lastQuotes string,
) (*Output, error) {

	quotes, err := uc.repo.ReadQuotes(lastQuotes)
	if err != nil {
		return nil, err
	}

	if len(quotes) == 0 {
		return nil, httpx.NotFound(
			"no quotes found",
			"there are no quotes available to generate metrics",
		)
	}

	carrierMetrics := make(map[string]*CarrierMetrics)

	var cheapest QuoteMetrics
	var expensive QuoteMetrics

	minPrice := math.MaxFloat64
	maxPrice := 0.0

	for _, q := range quotes {

		if _, exists := carrierMetrics[q.CarrierName]; !exists {
			carrierMetrics[q.CarrierName] = &CarrierMetrics{
				CarrierName: q.CarrierName,
			}
		}

		m := carrierMetrics[q.CarrierName]
		m.TotalQuotes++
		m.TotalPrice += q.Price

		if q.Price < minPrice {
			minPrice = q.Price
			cheapest = QuoteMetrics{
				CarrierName: q.CarrierName,
				Price:       q.Price,
			}
		}

		if q.Price > maxPrice {
			maxPrice = q.Price
			expensive = QuoteMetrics{
				CarrierName: q.CarrierName,
				Price:       q.Price,
			}
		}
	}

	var result []CarrierMetrics
	for _, m := range carrierMetrics {
		m.AveragePrice = m.TotalPrice / float64(m.TotalQuotes)
		result = append(result, *m)
	}

	return &Output{
		CarrierMetrics:   result,
		CheapestCarrier:  cheapest,
		ExpensiveCarrier: expensive,
	}, nil
}
