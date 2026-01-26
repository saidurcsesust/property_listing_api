package models

// PropertyResponse is used for API responses.
type PropertyResponse struct {
	Items []PropertyItem `json:"Items"`
}

type PropertyItem struct {
	ID        string   `json:"ID"`
	Feed      int      `json:"Feed"`
	Published bool     `json:"Published"`
	GeoInfo   GeoInfo  `json:"GeoInfo"`
	Property  Property `json:"Property"`
	Partner   Partner  `json:"Partner"`
}

type GeoInfo struct {
	Categories  []Category `json:"Categories"`
	City        string     `json:"City"`
	Country     string     `json:"Country"`
	CountryCode string     `json:"CountryCode"`
	Display     string     `json:"Display"`
	LocationID  string     `json:"LocationID"`
	Lat         string     `json:"Lat"`
	Lng         string     `json:"Lng"`
	Slug        string     `json:"Slug"`
}

type Category struct {
	Name    string   `json:"Name"`
	Slug    string   `json:"Slug"`
	Type    string   `json:"Type"`
	Display []string `json:"Display"`
}

type Property struct {
	Amenities              map[string]string `json:"Amenities"`
	Counts                 Counts            `json:"Counts"`
	FeatureImage           string            `json:"FeatureImage"`
	IsPetFriendly          bool              `json:"IsPetFriendly"`
	MinStay                int               `json:"MinStay"`
	PropertyName           string            `json:"PropertyName"`
	PropertySlug           string            `json:"PropertySlug"`
	PropertyType           string            `json:"PropertyType"`
	PropertyTypeCategoryId string            `json:"PropertyTypeCategoryId"`
	RoomSize               float64           `json:"RoomSize"`
}

type Counts struct {
	Bedroom   int `json:"Bedroom"`
	Bathroom  int `json:"Bathroom"`
	Occupancy int `json:"Occupancy"`
}

type Partner struct {
	ID           string   `json:"ID"`
	OwnerID      string   `json:"OwnerID"`
	Archived     []string `json:"Archived"`
	PropertyType string   `json:"PropertyType"`
	URL          string   `json:"URL"`
}

// RawPropertyDetail is used for upstream service responses.
type RawPropertyDetail struct {
	ID                   string        `json:"id"`
	Feed                 int           `json:"feed"`
	Published            bool          `json:"published"`
	Categories           string        `json:"categories"`
	City                 string        `json:"city"`
	Country              string        `json:"country"`
	CountryCode          string        `json:"country_code"`
	Display              string        `json:"display"`
	LocationID           string        `json:"location_id"`
	LonLat               LonLat        `json:"lonlat"`
	AmenityCategories    []string      `json:"amenity_categories"`
	BedroomCount         int           `json:"bedroom_count"`
	BathroomCount        int           `json:"bathroom_count"`
	Occupancy            int           `json:"occupancy"`
	FeatureImage         string        `json:"feature_image"`
	PropertyFlags        PropertyFlags `json:"property_flags"`
	MinStay              int           `json:"min_stay"`
	PropertyName         string        `json:"property_name"`
	PropertySlug         string        `json:"property_slug"`
	PropertyType         string        `json:"property_type"`
	RoomSizeSqft         float64       `json:"room_size_sqft"`
	OwnerID              *string       `json:"owner_id"`
	Archived             []string      `json:"archived"`
	FeedProviderID       string        `json:"feed_provider_id"`
	FeedProviderURL      string        `json:"feed_provider_url"`
	PropertyTypeCategory string        `json:"property_type_category"`
}

type LonLat struct {
	Coordinates []float64 `json:"coordinates"`
}

type PropertyFlags struct {
	IsPetFriendly bool `json:"is_pet_friendly"`
}
