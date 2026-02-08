package dtos

type QuoteRequest struct {
	Recipient Recipient `json:"recipient" validate:"required"`
	Volumes   []Volume  `json:"volumes" validate:"required,dive"`
}

type Recipient struct {
	Address Address `json:"address" validate:"required"`
}

type Address struct {
	Zipcode string `json:"zipcode" validate:"required"`
}

type Volume struct {
	Category      int64   `json:"category" validate:"required,gt=0"`
	Amount        float64 `json:"amount" validate:"required,gt=0"`
	UnitaryWeight int64   `json:"unitary_weight" validate:"required,gt=0"`
	Price         float64 `json:"price" validate:"required,gt=0"`
	Sku           string  `json:"sku" validate:"required"`
	Height        float64 `json:"height" validate:"required,gt=0"`
	Width         float64 `json:"width" validate:"required,gt=0"`
	Length        float64 `json:"length" validate:"required,gt=0"`
	UnitaryPrice  float64 `json:"unitary_price" validate:"required,gt=0"`
}
