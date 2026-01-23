package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"property_listing_api/models"
)

type PropertyService struct {
	baseURL string
	client  *HTTPClient
}

type propertyTypeInfo struct {
	name string
	id   string
}

var propertyTypeMappings = map[string]propertyTypeInfo{
	"House": {name: "Casa", id: "6"},
}

func NewPropertyService(baseURL string, client *HTTPClient) *PropertyService {
	return &PropertyService{baseURL: baseURL, client: client}
}

func (s *PropertyService) GetPropertyDetails(id string) (*models.RawPropertyDetail, error) {
	if strings.TrimSpace(s.baseURL) == "" {
		return nil, errors.New("property service base URL is empty")
	}
	url := fmt.Sprintf("%s/%s", strings.TrimRight(s.baseURL, "/"), id)

	var detail models.RawPropertyDetail
	_, err := s.client.GetJSON(url, &detail)
	if err != nil {
		return nil, err
	}

	return &detail, nil
}

func TransformPropertyDetail(detail *models.RawPropertyDetail) models.PropertyItem {
	categories := parseCategories(detail.Categories)
	lat, lng := formatCoordinates(detail.LonLat.Coordinates)
	slug := ""
	if len(categories) > 0 {
		slug = categories[len(categories)-1].Slug
	}

	amenities := mapAmenities(detail.AmenityCategories)

	propertyType := detail.PropertyTypeCategory
	propertyTypeCategoryID := ""
	if mapping, ok := propertyTypeMappings[detail.PropertyTypeCategory]; ok {
		propertyType = mapping.name
		propertyTypeCategoryID = mapping.id
	}

	ownerID := ""
	if detail.OwnerID != nil {
		ownerID = *detail.OwnerID
	}

	return models.PropertyItem{
		ID:        detail.ID,
		Feed:      detail.Feed,
		Published: detail.Published,
		GeoInfo: models.GeoInfo{
			Categories:  categories,
			City:        detail.City,
			Country:     detail.Country,
			CountryCode: detail.CountryCode,
			Display:     detail.Display,
			LocationID:  detail.LocationID,
			Lat:         lat,
			Lng:         lng,
			Slug:        slug,
		},
		Property: models.Property{
			Amenities: amenities,
			Counts: models.Counts{
				Bedroom:   detail.BedroomCount,
				Bathroom:  detail.BathroomCount,
				Occupancy: detail.Occupancy,
			},
			FeatureImage:           detail.FeatureImage,
			IsPetFriendly:          detail.PropertyFlags.IsPetFriendly,
			MinStay:                detail.MinStay,
			PropertyName:           detail.PropertyName,
			PropertySlug:           detail.PropertySlug,
			PropertyType:           propertyType,
			PropertyTypeCategoryId: propertyTypeCategoryID,
			RoomSize:               detail.RoomSizeSqft,
		},
		Partner: models.Partner{
			ID:           detail.FeedProviderID,
			OwnerID:      ownerID,
			Archived:     detail.Archived,
			PropertyType: detail.PropertyType,
			URL:          detail.FeedProviderURL,
		},
	}
}

func parseCategories(raw string) []models.Category {
	if strings.TrimSpace(raw) == "" {
		return nil
	}

	var categories []models.Category
	if err := json.Unmarshal([]byte(raw), &categories); err != nil {
		return nil
	}
	return categories
}

func formatCoordinates(coords []float64) (string, string) {
	if len(coords) < 2 {
		return "", ""
	}
	lng := strconv.FormatFloat(coords[0], 'f', -1, 64)
	lat := strconv.FormatFloat(coords[1], 'f', -1, 64)
	return lat, lng
}

func mapAmenities(amenityCategories []string) map[string]string {
	amenities := make(map[string]string)
	for index, name := range amenityCategories {
		if strings.TrimSpace(name) == "" {
			continue
		}
		key := strconv.Itoa(index + 1)
		amenities[key] = name
	}
	return amenities
}
