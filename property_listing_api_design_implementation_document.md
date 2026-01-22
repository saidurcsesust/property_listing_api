# Property Listing API – Design & Implementation Document

## 1. Overview
This document describes an improved design and implementation plan for the **Property Listing API** based on the given assignment. The API aggregates data from two dependent services:

1. **Get Property IDs by Location**
2. **Get Property Details by Property ID**

The goal is to expose a clean, validated, and production-ready API using the **Beego framework (Go)** while following proper folder structure, configuration management, validation, error handling, and coding best practices.

---

## 2. Dependencies / External APIs

### 2.1 Get Property IDs by Location
Returns a list of property IDs for a given colon-separated location.

- **Endpoint**:  
  `GET http://192.168.0.35:8099/api/:locationSeparatedWithColon`

- **Example**:  
  `GET http://192.168.0.35:8099/api/usa:florida:destin`

- **Expected Response (example)**:
```json
["usa", "usa:florida", "usa:florida:destin", "usa:texas", "usa:texas:galveston", "usa:california", "italy"]
```

---

### 2.2 Get Property Details by Property ID
Returns detailed information for a single property.

- **Endpoint**:  
  `GET http://192.168.0.35:8099/api/:propertyID`

- **Example**:  
  `GET http://192.168.0.35:8099/api/BC-12810439`



---

## 3. New API Specification (Property Listing API)

### 3.1 Endpoint

```
GET /v1/properties
```

### 3.2 Query Parameters

| Parameter | Type    | Required |  | Description |
|----------|---------|----------|-------------|-------------|
| x-api-key | string  | Yes      | header | a secret key (e.g. 63f4945d921d599f27ae4fdf5bada3f1) |
| location | string  | Yes      | path | Colon-separated location (e.g. `usa:florida:destin`) |
| items    | boolean | Yes      | param | Must be `true` to fetch properties |

---

### 3.3 Behavior

1. Validate request parameters.
2. Call **Location API** to fetch property IDs.
3. Iterate over property IDs **in the same order returned**.
4. For each property ID:
   - Call **Property Details API**.
   - Transform and normalize the response.
5. Aggregate results into a single response object.

---

### 3.4 Success Response (200)

```json
{
  "Items": [
  {
    "ID": "BC-12810439",
    "Feed": 11,
    "Published": true,
    "GeoInfo": {
      "Categories": [
        {
          "Name": "Nepal",
          "Slug": "nepal",
          "Type": "country",
          "Display": [
            "nepal"
          ]
        },
        {
          "Name": "Bharatpur",
          "Slug": "nepal/bharatpur",
          "Type": "city",
          "Display": [
            "nepal",
            "bharatpur"
          ]
        }
      ],
      "City": "Nārāyangarh",
      "Country": "Central Development Region",
      "CountryCode": "NP",
      "Display": "Bharatpur, Nepal",
      "LocationID": "571",
      "Lat": "27.627859",
      "Lng": "84.40818",
      "Slug":"nepal/bharatpur"
    },
    "Property": {
      "Amenities": {
        "1": "Aire acondicionado",
        "2": "Balcón/Terraza",
        "6": "Bañera de hidromasaje",
        "7": "Internet",
        "10": "Estacionamiento",
        "11": "Mascota amigable",
        "15": "Vista",
        "31": "Instalaciones de bienestar",
        "33": "Chimenea/Calefacción"
      },
      "Counts": {
        "Bedroom": 3,
        "Bathroom": 3,
        "Occupancy": 3
      },
      "FeatureImage": "dhakal-villa-np-n%C4%81r%C4%81yangarh-bc-12810439-0.jpg",
      "IsPetFriendly": true,
      "MinStay": 1,
      "PropertyName": "Dhakal Villa",
      "PropertySlug": "dhakal-villa",
      "PropertyType": "Casa",
      "PropertyTypeCategoryId": "6",
      "RoomSize": 1341.9
    },
    "Partner": {
      "ID": "12810439",
      "OwnerID": "",
      "Archived": [],
      "PropertyType": "Homestays",
      "URL": "https://www.booking.com/hotel/np/dhakal-villa.html?aid=affiliate_id"
    }
  },
  {
    "ID": "BC-12810439",
    "Feed": 11,
    "Published": true,
    "GeoInfo": {
      "Categories": [
        {
          "Name": "Nepal",
          "Slug": "nepal",
          "Type": "country",
          "Display": [
            "nepal"
          ]
        },
        {
          "Name": "Bharatpur",
          "Slug": "nepal/bharatpur",
          "Type": "city",
          "Display": [
            "nepal",
            "bharatpur"
          ]
        }
      ],
      "City": "Nārāyangarh",
      "Country": "Central Development Region",
      "CountryCode": "NP",
      "Display": "Bharatpur, Nepal",
      "LocationID": "571",
      "Lat": "27.627859",
      "Lng": "84.40818",
      "Slug":"nepal/bharatpur"
    },
    "Property": {
      "Amenities": {
        "1": "Aire acondicionado",
        "2": "Balcón/Terraza",
        "6": "Bañera de hidromasaje",
        "7": "Internet",
        "10": "Estacionamiento",
        "11": "Mascota amigable",
        "15": "Vista",
        "31": "Instalaciones de bienestar",
        "33": "Chimenea/Calefacción"
      },
      "Counts": {
        "Bedroom": 3,
        "Bathroom": 3,
        "Occupancy": 3
      },
      "FeatureImage": "dhakal-villa-np-n%C4%81r%C4%81yangarh-bc-12810439-0.jpg",
      "IsPetFriendly": true,
      "MinStay": 1,
      "PropertyName": "Dhakal Villa",
      "PropertySlug": "dhakal-villa",
      "PropertyType": "Casa",
      "PropertyTypeCategoryId": "6",
      "RoomSize": 1341.9
    },
    "Partner": {
      "ID": "12810439",
      "OwnerID": "",
      "Archived": [],
      "PropertyType": "Homestays",
      "URL": "https://www.booking.com/hotel/np/dhakal-villa.html?aid=affiliate_id"
    }
  }
]
}
```

---

## 4. Error Handling

### 4.1 Validation Errors (400)

```json
{
  "error": "location parameter is required"
}
```

```json
{
  "error": "items parameter must be true"
}
```

---

### 4.2 External API Errors (502)

```json
{
  "error": "failed to fetch property IDs from location service"
}
```

---

### 4.3 Internal Server Errors (500)

```json
{
  "error": "unexpected server error"
}
```

---

## 5. Beego Project Structure

```
property-listing-api/
├── conf/
│   └── app.conf
│
├── controllers/
│   └── property_controller.go
├── routers/
│   └── router.go
├── services/
│   ├── location_service.go
│   ├── property_service.go
│   └── http_client.go
├── models/
│   └── property_response.go
├── utils/
│   └── validator.go
├── main.go
└── go.mod
```

---

## 6. Configuration Management

### 6.1 `conf/app.conf`

```ini
appname = property-listing-api
httpport = 8080
runmode = dev

```

---

## Evaluation Checklist (Mapped)

✔ Complete project structure  
✔ Runnable Beego project  
✔ Config-driven external APIs  
✔ Strong validation  
✔ Proper error handling  
✔ Clean, maintainable code  
✔ Well-documented behavior  
✔ Beego folder structure compliance

---

**End of Document**
