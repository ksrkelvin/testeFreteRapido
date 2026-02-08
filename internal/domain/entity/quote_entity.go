package entity

type QuoteEntity struct {
	CarrierName  string
	Service      string
	DeadlineDays int
	Price        float64
}
type VolumeEntity struct {
	Category      int64
	Amount        float64
	UnitaryWeight int64
	Price         float64
	Sku           string
	Height        float64
	Width         float64
	Length        float64
	UnitaryPrice  float64
}
