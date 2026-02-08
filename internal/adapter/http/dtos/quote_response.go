package dtos

type QuoteResponse struct {
	Carrier []Carrier `json:"carrier"`
}

type Carrier struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline string  `json:"deadline"`
	Price    float64 `json:"price"`
}
