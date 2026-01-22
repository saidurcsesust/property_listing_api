package utils

import (
	"errors"
	"strings"
)

func ValidateQueryParams(location, items string) error {
	if strings.TrimSpace(location) == "" {
		return errors.New("location parameter is required")
	}
	if items != "true" {
		return errors.New("items parameter must be true")
	}
	return nil
}

func ValidateAPIKey(headerValue, expected string) error {
	if strings.TrimSpace(expected) == "" {
		return errors.New("api key is not configured")
	}
	if strings.TrimSpace(headerValue) == "" {
		return errors.New("x-api-key header is required")
	}
	if headerValue != expected {
		return errors.New("invalid x-api-key")
	}
	return nil
}
