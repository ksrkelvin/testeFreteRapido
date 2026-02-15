package freteRapidoApi_test

import (
	"context"
	"reflect"
	"testeFreteRapido/pkg/freteRapidoApi"
	"testing"
)

func TestClient_QuoteV3(t *testing.T) {
	tests := []struct {
		name        string
		mockResp    []byte
		mockStatus  int
		mockErr     error
		recipient   freteRapidoApi.Recipient
		dispatchers []freteRapidoApi.DispatcherRequest
		want        freteRapidoApi.ResponseCotacaoFreteV3
		wantErr     bool
	}{
		{
			name: "success with multiple volumes",
			mockResp: []byte(`{
				"dispatchers": [
					{
						"id": "disp_1",
						"request_id": "req_1",
						"registered_number_shipper": "123456789",
						"registered_number_dispatcher": "987654321",
						"zipcode_origin": 12345678,
						"offers": [
							{
								"offer": 1,
								"simulation_type": 0,
								"carrier": {
									"name": "Correios",
									"registered_number": "111222333",
									"state_inscription": "123456",
									"logo": "url_logo",
									"reference": 1,
									"company_name": "Correios S.A."
								},
								"service": "Sedex",
								"service_code": "SEDEX",
								"delivery_time": { "days": 3, "estimated_date": "2026-02-20" },
								"expiration": "2026-02-21",
								"cost_price": 50.0,
								"final_price": 60.0,
								"weights": { "real": 13 },
								"original_delivery_time": { "days": 3, "estimated_date": "2026-02-20" },
								"identifier": "offer_1",
								"home_delivery": true,
								"carrier_original_delivery_time": { "days": 3, "estimated_date": "2026-02-20" },
								"modal": "road",
								"esg": { "co2_emission_estimate": 2.5 }
							}
						]
					}
				]
			}`),
			mockStatus: 200,
			mockErr:    nil,
			recipient: freteRapidoApi.Recipient{
				Zipcode: 01311000,
			},
			dispatchers: []freteRapidoApi.DispatcherRequest{
				{
					RegisteredNumber: "987654321",
					Zipcode:          01311000,
					Volumes: []freteRapidoApi.Volume{
						{
							Category:      "7",
							Amount:        1,
							UnitaryWeight: 5,
							Price:         349,
							Sku:           "abc-teste-123",
							Height:        0.2,
							Width:         0.2,
							Length:        0.2,
							UnitaryPrice:  10,
						},
						{
							Category:      "7",
							Amount:        2,
							UnitaryWeight: 4,
							Price:         556,
							Sku:           "abc-teste-527",
							Height:        0.4,
							Width:         0.6,
							Length:        0.15,
							UnitaryPrice:  10,
						},
					},
				},
			},
			want: freteRapidoApi.ResponseCotacaoFreteV3{
				Dispatchers: []freteRapidoApi.DispatcherResponse{
					{
						ID:                         "disp_1",
						RequestID:                  "req_1",
						RegisteredNumberShipper:    "123456789",
						RegisteredNumberDispatcher: "987654321",
						ZipcodeOrigin:              12345678,
						Offers: []freteRapidoApi.Offer{
							{
								Offer:          1,
								SimulationType: 0,
								Carrier: freteRapidoApi.Carrier{
									Name:             "Correios",
									RegisteredNumber: "111222333",
									StateInscription: "123456",
									Logo:             "url_logo",
									Reference:        1,
									CompanyName:      "Correios S.A.",
								},
								Service:                     "Sedex",
								ServiceCode:                 "SEDEX",
								DeliveryTime:                freteRapidoApi.DeliveryTime{Days: 3, EstimatedDate: "2026-02-20"},
								Expiration:                  "2026-02-21",
								CostPrice:                   50.0,
								FinalPrice:                  60.0,
								Weights:                     freteRapidoApi.Weights{Real: 13},
								OriginalDeliveryTime:        freteRapidoApi.DeliveryTime{Days: 3, EstimatedDate: "2026-02-20"},
								Identifier:                  "offer_1",
								HomeDelivery:                true,
								CarrierOriginalDeliveryTime: freteRapidoApi.DeliveryTime{Days: 3, EstimatedDate: "2026-02-20"},
								Modal:                       "road",
								Esg:                         freteRapidoApi.Esg{Co2EmissionEstimate: 2.5},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := newTestClient(tt.mockResp, tt.mockStatus, tt.mockErr)

			got, err := client.QuoteV3(context.Background(), tt.recipient, tt.dispatchers)
			if (err != nil) != tt.wantErr {
				t.Fatalf("QuoteV3() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("QuoteV3() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
