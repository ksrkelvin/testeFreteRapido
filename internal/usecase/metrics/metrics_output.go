package metrics

type Output struct {
	CarrierMetrics   []CarrierMetrics
	CheapestCarrier  QuoteMetrics
	ExpensiveCarrier QuoteMetrics
}

type CarrierMetrics struct {
	CarrierName  string
	TotalQuotes  int
	TotalPrice   float64
	AveragePrice float64
}

type QuoteMetrics struct {
	CarrierName string
	Price       float64
}
