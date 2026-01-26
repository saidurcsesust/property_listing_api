# Property Listing API

A small Beego-based service that aggregates property IDs by location and returns
normalized property details.

## Requirements

- Go 1.25+ (per `go.mod`)
- Access to the upstream location and property services

## Configuration

Create a `.env` in the project root:

```
APP_NAME=property_listing_api
HTTP_PORT=8080
RUN_MODE=dev
LOCATION_SERVICE_URL=http://192.168.0.35:8099/api
PROPERTY_SERVICE_URL=http://192.168.0.35:8099/api
API_KEY=your_api_key_here
HTTP_CLIENT_TIMEOUT_SECONDS=10
```

These values are loaded at startup and applied to Beego config keys.

## Run

```
go run .
```

The service listens on `http://localhost:8080` by default.

## API

### Get properties

```
GET /v1/properties?location=usa:florida:destin&items=true
```

Headers:

```
x-api-key: <API_KEY>
```

Response:

- `200`: `models.PropertyResponse`
- `400`: invalid query params
- `401`: invalid API key
- `502`: upstream service failure
- `500`: server configuration error

## Development notes

- `conf/app.conf` maps to Beego settings; env values override these at runtime.
- The homepage is at `/` and includes a small “Try it now” fetch UI.
