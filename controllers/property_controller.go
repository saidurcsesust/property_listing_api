package controllers

import (
	"net/http"
	"time"

	"property_listing_api/models"
	"property_listing_api/services"
	"property_listing_api/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	beego.Controller
}

func (c *PropertyController) Get() {
	apiKey, err := beego.AppConfig.String("api_key")
	if err != nil || apiKey == "" {
		c.writeError(http.StatusInternalServerError, "unexpected server error")
		return
	}

	if err := utils.ValidateAPIKey(c.Ctx.Input.Header("x-api-key"), apiKey); err != nil {
		c.writeError(http.StatusBadRequest, err.Error())
		return
	}

	location := c.GetString("location")
	itemsParam := c.GetString("items")
	if err := utils.ValidateQueryParams(location, itemsParam); err != nil {
		c.writeError(http.StatusBadRequest, err.Error())
		return
	}

	locationBaseURL, err := beego.AppConfig.String("location_service_url")
	if err != nil {
		c.writeError(http.StatusInternalServerError, "unexpected server error")
		return
	}
	propertyBaseURL, err := beego.AppConfig.String("property_service_url")
	if err != nil {
		c.writeError(http.StatusInternalServerError, "unexpected server error")
		return
	}
	if locationBaseURL == "" || propertyBaseURL == "" {
		c.writeError(http.StatusInternalServerError, "unexpected server error")
		return
	}

	timeoutSeconds, err := beego.AppConfig.Int("http_client_timeout_seconds")
	if err != nil || timeoutSeconds <= 0 {
		timeoutSeconds = 10
	}

	client := services.NewHTTPClient(time.Duration(timeoutSeconds) * time.Second)
	locationService := services.NewLocationService(locationBaseURL, client)
	propertyService := services.NewPropertyService(propertyBaseURL, client)

	ids, err := locationService.GetPropertyIDs(location)
	if err != nil {
		c.writeError(http.StatusBadGateway, "failed to fetch property IDs from location service")
		return
	}

	items := make([]models.PropertyItem, 0, len(ids))
	for _, id := range ids {
		detail, err := propertyService.GetPropertyDetails(id)
		if err != nil {
			c.writeError(http.StatusBadGateway, "failed to fetch property details from property service")
			return
		}
		items = append(items, services.TransformPropertyDetail(detail))
	}

	c.Data["json"] = models.PropertyResponse{Items: items}
	c.ServeJSON()
}

func (c *PropertyController) writeError(status int, message string) {
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = map[string]string{"error": message}
	c.ServeJSON()
}
