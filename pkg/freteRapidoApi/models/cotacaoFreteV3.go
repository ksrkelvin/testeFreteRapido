package models

type RequestCotacaoFreteV3 struct {
	Shipper        Shipper             `json:"shipper"`
	Recipient      Recipient           `json:"recipient"`
	Dispatchers    []DispatcherRequest `json:"dispatchers"`
	SimulationType []int               `json:"simulation_type"`
}

type ResponseCotacaoFreteV3 struct {
	Dispatchers []DispatcherResponse `json:"dispatchers"`
}

type DispatcherRequest struct {
	RegisteredNumber string   `json:"registered_number"`
	Zipcode          int32    `json:"zipcode"`
	Volumes          []Volume `json:"volumes"`
}

type DispatcherResponse struct {
	ID                         string  `json:"id"`
	RequestID                  string  `json:"request_id"`
	RegisteredNumberShipper    string  `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string  `json:"registered_number_dispatcher"`
	ZipcodeOrigin              int32   `json:"zipcode_origin"`
	Offers                     []Offer `json:"offers"`
}

type Offer struct {
	Offer                       int64        `json:"offer"`
	SimulationType              int64        `json:"simulation_type"`
	Carrier                     Carrier      `json:"carrier"`
	Service                     string       `json:"service"`
	ServiceCode                 string       `json:"service_code"`
	DeliveryTime                DeliveryTime `json:"delivery_time"`
	Expiration                  string       `json:"expiration"`
	CostPrice                   float64      `json:"cost_price"`
	FinalPrice                  float64      `json:"final_price"`
	Weights                     Weights      `json:"weights"`
	OriginalDeliveryTime        DeliveryTime `json:"original_delivery_time"`
	Identifier                  string       `json:"identifier"`
	HomeDelivery                bool         `json:"home_delivery"`
	CarrierOriginalDeliveryTime DeliveryTime `json:"carrier_original_delivery_time"`
	Modal                       string       `json:"modal"`
	Esg                         Esg          `json:"esg"`
}

type Carrier struct {
	Name             string `json:"name"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Logo             string `json:"logo"`
	Reference        int64  `json:"reference"`
	CompanyName      string `json:"company_name"`
}

type DeliveryTime struct {
	Days          int64  `json:"days"`
	EstimatedDate string `json:"estimated_date"`
}

type Esg struct {
	Co2EmissionEstimate float64 `json:"co2_emission_estimate"`
}

type Weights struct {
	Real int64 `json:"real"`
}

type Volume struct {
	Category      string  `json:"category"`
	Amount        int64   `json:"amount"`
	UnitaryWeight int64   `json:"unitary_weight"`
	Price         int64   `json:"price"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
	UnitaryPrice  float64 `json:"unitary_price"`
}

type Recipient struct {
	Zipcode int64 `json:"zipcode"`
}

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}
