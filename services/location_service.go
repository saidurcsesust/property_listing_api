package services

import (
	"errors"
	"fmt"
	"strings"
)

type LocationService struct {
	baseURL string
	client  *HTTPClient
}

func NewLocationService(baseURL string, client *HTTPClient) *LocationService {
	return &LocationService{baseURL: baseURL, client: client}
}

func (s *LocationService) GetPropertyIDs(location string) ([]string, error) {
	if strings.TrimSpace(s.baseURL) == "" {
		return nil, errors.New("location service base URL is empty")
	}
	url := fmt.Sprintf("%s/%s", strings.TrimRight(s.baseURL, "/"), location)

	var ids []string
	_, err := s.client.GetJSON(url, &ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
