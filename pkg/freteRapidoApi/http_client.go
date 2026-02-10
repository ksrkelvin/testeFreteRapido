package freteRapidoApi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"
)

type httpClient struct {
	baseURL string
	client  *http.Client
}

func newHTTPClient(baseURL string, timeout time.Duration) *httpClient {
	return &httpClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (h *httpClient) Post(
	ctx context.Context,
	path string,
	headers map[string]string,
	body []byte,
) ([]byte, int, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		h.baseURL+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, 0, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := h.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	return respBody, res.StatusCode, err
}
