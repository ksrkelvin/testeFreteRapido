package interfaces

import "context"

type Volume struct {
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

type FreightQuote struct {
	CarrierName  string
	Service      string
	DeadlineDays int64
	FinalPrice   float64
}

type FreightGateway interface {
	Quote(ctx context.Context, recipientZip int32, volumes []Volume) ([]FreightQuote, error)
}
