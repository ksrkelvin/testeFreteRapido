package dtos

type QuoteResponseDTO struct {
	Carrier []CarrierDTO `json:"carrier"`
}

type CarrierDTO struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline string  `json:"deadline"`
	Price    float64 `json:"price"`
}
