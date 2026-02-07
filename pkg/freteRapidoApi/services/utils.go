package services

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
)

type utils struct {
	Host string
}

type Utils interface {
	ReqHTTP(method string, headers map[string]string, body []byte) (response string, statusHttp int, err error)
}

func NewUtils(host string) *utils {
	return &utils{
		Host: host,
	}
}

func (u *utils) ReqHTTP(path string, method string, headers map[string]string, requestBody []byte) (responseBody []byte, statusHttp int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(
				"Error to try ReqHTTP: %v\n\nstack trace:\n%s",
				r,
				debug.Stack(),
			)
		}
	}()

	req, err := http.NewRequest(
		method,
		u.Host+path,
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return nil, 0, fmt.Errorf("error creating request: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("Error to try make HTTP request: %v", err)
	}
	responseBody, err = io.ReadAll(response.Body)
	defer response.Body.Close()

	fmt.Println(string(requestBody))
	fmt.Printf("Response body: %s\n", string(responseBody))

	return responseBody, response.StatusCode, nil
}
