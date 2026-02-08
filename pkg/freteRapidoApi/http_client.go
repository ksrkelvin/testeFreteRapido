package freteRapidoApi

import (
	"bytes"
	"io"
	"net/http"
)

type httpClient struct {
	host string
}

func newHTTPClient(host string) *httpClient {
	return &httpClient{host: host}
}

func (h *httpClient) Post(path string, headers map[string]string, body []byte) ([]byte, int, error) {
	req, err := http.NewRequest(http.MethodPost, h.host+path, bytes.NewReader(body))
	if err != nil {
		return nil, 0, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	return respBody, res.StatusCode, err
}
