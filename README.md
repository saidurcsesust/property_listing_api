# Property Listing API

A small Beego-based service that aggregates property IDs by location and returns
normalized property details.

## Highlights

- Versioned API under `/v1`
- Location lookup + property aggregation
- API key protected endpoints
- Normalized response payloads

## Tech stack

- Go + Beego (v2)
- HTTP clients for upstream services

## Requirements

- Go 1.25+ (per `go.mod`)

## Clone

```
git clone https://github.com/saidurcsesust/property_listing_api.git
cd property_listing_api
```

## Run

```
bee run
```

The service listens on `http://localhost:8080` by default.

## Project structure

```
.
├── conf/
│   └── app.conf                   # Beego configuration defaults
├── controllers/
│   ├── default.go                 # Root route handler
│   └── property_controller.go     # /v1/properties handler
├── models/
│   └── property_response.go       # API response models
├── routers/
│   └── router.go                  # Route and namespace setup
├── services/
│   ├── http_client.go             # HTTP client wrapper
│   ├── location_service.go        # Location lookup/normalization
│   └── property_service.go        # Property aggregation and mapping
├── static/
│   └── js/
│       └── reload.min.js          # Beego dev reload script
├── tests/
│   └── default_test.go            # Basic Beego test
├── utils/
│   └── validator.go               # Request validation helpers
├── views/
│   └── index.tpl                  # Default template
├── main.go                        # Application bootstrap
├── go.mod                         # Module definition
└── go.sum                         # Dependency checksums
```

## Request flow

```
router -> controller -> service(s) -> model -> JSON response
```

## API

### Get properties

```
GET /v1/properties/usa:florida:destin?items=true
```

Headers:

```
x-api-key: 63f4945d921d599f27ae4fdf5bada3f1
```

Example curl:

```bash
curl -s \
  -H "x-api-key: 63f4945d921d599f27ae4fdf5bada3f1" \
  "http://localhost:8080/v1/properties/usa:florida:destin?items=true"
```

Response:

- `200`: `models.PropertyResponse`
- `400`: invalid query params
- `401`: invalid API key
- `500`: server configuration error
- `502`: external API error


Example response (single item):

```json
{
  "Items": [
    {
      "ID": "HA-321392925",
      "Feed": 12,
      "Published": true,
      "GeoInfo": {
        "Categories": [
          {
            "Name": "USA",
            "Slug": "usa",
            "Type": "country",
            "Display": ["usa"]
          },
          {
            "Name": "Virginia",
            "Slug": "usa/virginia",
            "Type": "state",
            "Display": ["usa", "virginia"]
          },
          {
            "Name": "Roanoke",
            "Slug": "usa/virginia/roanoke",
            "Type": "city",
            "Display": ["usa", "virginia", "roanoke"]
          },
          {
            "Name": "Montvale",
            "Slug": "usa/virginia/montvale",
            "Type": "city",
            "Display": ["usa", "virginia", "montvale"]
          }
        ],
        "City": "Blue Ridge",
        "Country": "USA",
        "CountryCode": "US",
        "Display": "Blue Ridge, Virginia, United States",
        "LocationID": "6348039",
        "Lat": "37.378845",
        "Lng": "-79.731461",
        "Slug": "usa/virginia/montvale"
      },
      "Property": {
        "Amenities": {
          "1": "Air Conditioner",
          "10": "Wellness Facilities",
          "11": "Fireplace/Heating",
          "12": "Entertainment",
          "13": "Barbecue/Outdoor Cooking",
          "2": "Balcony/Terrace",
          "3": "Bedding/Linens",
          "4": "Hot Tub",
          "5": "Kitchen",
          "6": "Laundry",
          "7": "Parking",
          "8": "TV",
          "9": "Security/Safety"
        },
        "Counts": {
          "Bedroom": 1,
          "Bathroom": 1,
          "Occupancy": 2
        },
        "FeatureImage": "couples-getaway-with-us-blue-ridge-ha-321392925-0.jpg",
        "IsPetFriendly": false,
        "MinStay": 1,
        "PropertyName": "Couples Getaway with a Great Mountain View",
        "PropertySlug": "couples-getaway-with-a-great-mountain-view",
        "PropertyType": "Cabin",
        "PropertyTypeCategoryId": "",
        "RoomSize": 0
      },
      "Partner": {
        "ID": "392925",
        "OwnerID": "33960608",
        "Archived": ["VRBO", "EP"],
        "PropertyType": "Cabin",
        "URL": "https://www.vrbo.com/search?selected=33960608&regionId=6348039"
      }
    }
  ]
}
```

## Configuration

`conf/app.conf` provides defaults; environment variables override at runtime.

| Key                         | Purpose                              | Default |
| --------------------------- | ------------------------------------ | ------- |
| `appname`                   | Beego app name                       | `property_listing_api` |
| `httpport`                  | HTTP port                            | `8080` |
| `runmode`                   | Beego run mode                       | `dev` |
| `location_service_url`      | Location lookup service              | `http://192.168.0.35:8099/api` |
| `property_service_url`      | Property detail service              | `http://192.168.0.35:8099/api` |
| `api_key`                   | API key for incoming requests        | `63f4945d921d599f27ae4fdf5bada3f1` |
| `http_client_timeout_seconds` | Upstream HTTP timeout (seconds)    | `10` |

`x-api-key` is required for protected endpoints.
