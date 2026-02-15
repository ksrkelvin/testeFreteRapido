package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"testeFreteRapido/internal/adapter/http/handlers"
	"testeFreteRapido/internal/usecase/quote"
)

func TestQuoteHandler_Quote(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		body       interface{}
		mock       *MockQuoteUsecase
		wantStatus int
		wantError  bool
	}{
		{
			name: "invalid json body",
			body: "invalid",
			mock: &MockQuoteUsecase{
				ExecuteFunc: func(ctx context.Context, input quote.Input) (*quote.Output, error) {
					t.Fatal("usecase should not be called")
					return nil, nil
				},
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
		{
			name: "usecase returns error",
			body: map[string]interface{}{
				"recipient": map[string]interface{}{
					"address": map[string]interface{}{
						"zipcode": "01311000",
					},
				},
				"volumes": []map[string]interface{}{
					{
						"category":       7,
						"amount":         1,
						"unitary_weight": 5,
						"price":          349,
						"sku":            "abc",
						"height":         0.2,
						"width":          0.2,
						"length":         0.2,
						"unitary_price":  10,
					},
				},
			},
			mock: &MockQuoteUsecase{
				ExecuteFunc: func(ctx context.Context, input quote.Input) (*quote.Output, error) {
					return nil, errors.New("internal error")
				},
			},
			wantStatus: http.StatusInternalServerError,
			wantError:  true,
		},
		{
			name: "success",
			body: map[string]interface{}{
				"recipient": map[string]interface{}{
					"address": map[string]interface{}{
						"zipcode": "01311000",
					},
				},
				"volumes": []map[string]interface{}{
					{
						"category":       7,
						"amount":         1,
						"unitary_weight": 5,
						"price":          349,
						"sku":            "abc-teste-123",
						"height":         0.2,
						"width":          0.2,
						"length":         0.2,
						"unitary_price":  10,
					},
				},
			},
			mock: &MockQuoteUsecase{
				ExecuteFunc: func(ctx context.Context, input quote.Input) (*quote.Output, error) {
					if input.RecipientZip != 1311000 {
						t.Fatalf("unexpected zipcode: %v", input.RecipientZip)
					}
					return &quote.Output{}, nil
				},
			},
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name: "invalid request data - mapper error",
			body: map[string]interface{}{
				"recipient": map[string]interface{}{
					"address": map[string]interface{}{
						"zipcode": "",
					},
				},
				"volumes": []map[string]interface{}{},
			},
			mock: &MockQuoteUsecase{
				ExecuteFunc: func(ctx context.Context, input quote.Input) (*quote.Output, error) {
					t.Fatal("usecase should not be called")
					return nil, nil
				},
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			handler := handlers.NewQuoteHandler(tt.mock)

			router := gin.New()
			router.POST("/quote", handler.Quote)

			jsonBody, _ := json.Marshal(tt.body)

			req := httptest.NewRequest(
				http.MethodPost,
				"/quote",
				bytes.NewBuffer(jsonBody),
			)
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Fatalf("status = %v, want %v", w.Code, tt.wantStatus)
			}

			var response map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("invalid json response: %v", err)
			}

			if tt.wantError {
				if _, ok := response["message"]; !ok {
					t.Errorf("expected 'message' field in error response")
				}
			} else {
				if _, ok := response["carrier"]; !ok {
					t.Errorf("expected 'carriers' field in success response")
				}
			}

		})
	}
}
