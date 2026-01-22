package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	client *http.Client
}

func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{Timeout: timeout},
	}
}

func (h *HTTPClient) GetJSON(url string, target interface{}) (int, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}

	response, err := h.client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		_, _ = io.Copy(io.Discard, response.Body)
		return response.StatusCode, fmt.Errorf("unexpected status code %d", response.StatusCode)
	}

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(target); err != nil {
		return response.StatusCode, err
	}

	return response.StatusCode, nil
}
