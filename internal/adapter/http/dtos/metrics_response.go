package dtos

type MetricsResponse struct {
	CarrierMetrics   []CarrierMetrics `json:"carrier_metrics"`
	CheapestCarrier  QuoteMetrics     `json:"cheapest_carrier"`
	ExpensiveCarrier QuoteMetrics     `json:"expensive_carrier"`
}

type CarrierMetrics struct {
	CarrierName  string  `json:"carrier_name"`
	TotalQuotes  int     `json:"total_quotes"`
	TotalPrice   float64 `json:"total_price"`
	AveragePrice float64 `json:"average_price"`
}

type QuoteMetrics struct {
	CarrierName string  `json:"carrier_name"`
	Price       float64 `json:"price"`
}
