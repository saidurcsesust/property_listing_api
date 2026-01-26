# Property Listing API

A small Beego-based service that aggregates property IDs by location and returns
normalized property details.

## Requirements

- Go 1.25+ (per `go.mod`)

## Configuration

Create a `.env` in the project root:

```
LOCATION_SERVICE_URL= Location_Url
PROPERTY_SERVICE_URL= Property_Url
API_KEY=your_api_key_here
```

These values are loaded at startup and applied to Beego config keys.

## Run

```
bee run
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
