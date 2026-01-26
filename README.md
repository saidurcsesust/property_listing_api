# Property Listing API

A small Beego-based service that aggregates property IDs by location and returns
normalized property details.

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

## API

### Get properties

```
GET /v1/properties?location=usa:florida:destin&items=true
```

Headers:

```
x-api-key: 63f4945d921d599f27ae4fdf5bada3f1
```

Response:

- `200`: `models.PropertyResponse`
- `400`: invalid query params
- `401`: invalid API key
- `500`: server configuration error

## Development notes

- `conf/app.conf` maps to Beego settings; env values override these at runtime.
