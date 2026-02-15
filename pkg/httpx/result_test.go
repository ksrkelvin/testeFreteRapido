package httpx_test

import (
	"errors"
	"net/http"
	"reflect"
	"testeFreteRapido/pkg/httpx"
	"testing"
)

func TestFromResult(t *testing.T) {
	tests := []struct {
		name     string
		data     any
		err      error
		wantCode int
		wantResp any
	}{
		{
			name:     "Data is not nil, no error",
			data:     "success",
			err:      nil,
			wantCode: http.StatusOK,
			wantResp: "success",
		},
		{
			name:     "Data is nil, no error",
			data:     nil,
			err:      nil,
			wantCode: http.StatusNoContent,
			wantResp: nil,
		},
		{
			name:     "Error is not AppError",
			data:     nil,
			err:      errors.New("some error"),
			wantCode: http.StatusInternalServerError,
			wantResp: httpx.ErrorResponse{
				Message:     "internal server error",
				Description: "some error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCode, gotResp := httpx.FromResult(tt.data, tt.err)
			if gotCode != tt.wantCode {
				t.Errorf("FromResult() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("FromResult() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
