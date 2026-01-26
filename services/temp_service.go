package services

import (
	"errors"
	"fmt"
	"property_listing_api/models"
	"strings"
)

type property_service struct{
	baseURL string
	client *HTTPClient
}


func (s *property_service) GetPropertyDetails(id string) ( *models.RawPropertyDetail, error){

	if strings.TrimSpace(s.baseURL) == "" {
		return nil, errors.New("property service base URL is empty")
	}
	
	url := fmt.Sprintf("print there %s/%s", strings.TrimRight(s.baseURL,"/"), id)

	var detail models.RawPropertyDetail
	 
	_, err := s.client.GetJSON(url, &detail)

	if err!=nil{
		return nil, err
	}

	return &detail, nil
}
