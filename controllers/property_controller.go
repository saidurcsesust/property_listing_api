package controllers

import (
	"net/http"
	"sync"
	"time"

	"property_listing_api/models"
	"property_listing_api/services"
	"property_listing_api/utils"

	"github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	web.Controller
}

// @Title GetProperties
// @Description Get property listings by location
// @Param x-api-key header string true "API key"
// @Param location query string true "Location identifier"
// @Param items query string true "Number of items to return"
// @Success 200 {object} models.PropertyResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 502 {object} map[string]string
// @Failure 500 {object} map[string]string
// @router / [get]

func (c *PropertyController) Get() {

	// read api key from app.conf
	apiKey, err := web.AppConfig.String("api_key")
	if err != nil || apiKey == "" {
		c.writeError(http.StatusInternalServerError, "unexpected server error")
		return
	}

	// validate x-api-key header from client
	if err := utils.ValidateAPIKey(c.Ctx.Input.Header("x-api-key"), apiKey); err != nil {
		c.writeError(http.StatusUnauthorized, err.Error())
		return
	}

	// read and validate query parameters
	location := c.GetString("location")
	itemsParam := c.GetString("items")
	if err := utils.ValidateQueryParams(location, itemsParam); err != nil {
		c.writeError(http.StatusBadRequest, err.Error())
		return
	}

	locationBaseURL, err := web.AppConfig.String("location_service_url")
	if err != nil {
		c.writeError(http.StatusInternalServerError, "unexpected server error")
		return
	}
	propertyBaseURL, err := web.AppConfig.String("property_service_url")
	if err != nil {
		c.writeError(http.StatusInternalServerError, "unexpected server error")
		return
	}
	if locationBaseURL == "" || propertyBaseURL == "" {
		c.writeError(http.StatusInternalServerError, "unexpected server error")
		return
	}

	timeoutSeconds, err := web.AppConfig.Int("http_client_timeout_seconds")
	if err != nil || timeoutSeconds <= 0 {
		timeoutSeconds = 10
	}

	client := services.NewHTTPClient(time.Duration(timeoutSeconds) * time.Second)
	locationService := services.NewLocationService(locationBaseURL, client)
	propertyService := services.NewPropertyService(propertyBaseURL, client)

	// get ids
	ids, err := locationService.GetPropertyIDs(location)
	if err != nil {
		c.writeError(http.StatusBadGateway, "failed to fetch property IDs from location service")
		return
	}

	type propertyDetailResult struct {
		index  int
		detail *models.RawPropertyDetail
		err    error
	}

	results := make(chan propertyDetailResult, len(ids))
	var wg sync.WaitGroup

	// call property service for details concurrently
	wg.Add(len(ids))
	for index, id := range ids {
		go func(idx int, propertyID string) {
			defer wg.Done()
			detail, err := propertyService.GetPropertyDetails(propertyID)
			results <- propertyDetailResult{index: idx, detail: detail, err: err}
		}(index, id)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	items := make([]models.PropertyItem, len(ids))
	var fetchErr error
	for result := range results {
		if result.err != nil {
			if fetchErr == nil {
				fetchErr = result.err
			}
			continue
		}
		items[result.index] = services.TransformPropertyDetail(result.detail)
	}

	if fetchErr != nil {
		c.writeError(http.StatusBadGateway, "failed to fetch property details from property service")
		return
	}

	c.Data["json"] = models.PropertyResponse{Items: items}
	c.ServeJSON()
}

func (c *PropertyController) writeError(status int, message string) {
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = map[string]string{"error": message}
	c.ServeJSON()
}
