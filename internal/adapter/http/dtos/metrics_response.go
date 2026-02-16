package dtos

type MetricsResponseDTO struct {
	CarrierMetrics   []CarrierMetricsDTO `json:"carrier_metrics"`
	CheapestCarrier  QuoteMetricsDTO     `json:"cheapest_carrier"`
	ExpensiveCarrier QuoteMetricsDTO     `json:"expensive_carrier"`
}

type CarrierMetricsDTO struct {
	CarrierName  string  `json:"carrier_name"`
	TotalQuotes  int     `json:"total_quotes"`
	TotalPrice   float64 `json:"total_price"`
	AveragePrice float64 `json:"average_price"`
}

type QuoteMetricsDTO struct {
	CarrierName string  `json:"carrier_name"`
	Price       float64 `json:"price"`
}
