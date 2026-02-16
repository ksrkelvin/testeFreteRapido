package handlers_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testeFreteRapido/internal/adapter/http/handlers"
	"testeFreteRapido/internal/usecase/metrics"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMetricsHandler_Metrics(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		queryParam string
		usecase    *MockMetricsUsecase
		wantStatus int
		wantErr    bool
	}{
		{
			name:       "usecase returns error",
			queryParam: "10",
			usecase: &MockMetricsUsecase{
				ExecuteFunc: func(ctx context.Context, lastQuotes string) (*metrics.Output, error) {
					return nil, errors.New("error")
				},
			},
			wantStatus: http.StatusInternalServerError,
			wantErr:    true,
		},
		{
			name:       "success",
			queryParam: "10",
			usecase: &MockMetricsUsecase{
				ExecuteFunc: func(ctx context.Context, lastQuotes string) (*metrics.Output, error) {
					return &metrics.Output{}, nil
				},
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			handler := handlers.NewMetricsHandler(tt.usecase)

			router := gin.New()
			router.GET("/metrics", handler.Metrics)

			req := httptest.NewRequest(
				http.MethodGet,
				"/metrics?last_quotes="+tt.queryParam,
				nil,
			)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("status = %v, want %v", w.Code, tt.wantStatus)
			}

			if (w.Code != http.StatusOK) != tt.wantErr {
				t.Errorf("error status mismatch")
			}
		})
	}
}
