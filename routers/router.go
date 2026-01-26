package routers

import (
	"encoding/json"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"property_listing_api/controllers"
	"property_listing_api/utils"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		rw.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(rw).Encode(map[string]string{"error": "invalid route"})
	})

	// Validate API key for all /v1 routes.
	beego.InsertFilter("/v1/*", beego.BeforeRouter, func(ctx *context.Context) {
		apiKey, err := beego.AppConfig.String("api_key")
		if err != nil || apiKey == "" {
			writeJSONError(ctx, 500, "unexpected server error")
			return
		}
		if err := utils.ValidateAPIKey(ctx.Input.Header("x-api-key"), apiKey); err != nil {
			writeJSONError(ctx, 401, err.Error())
			return
		}
	})

	apiV1 := beego.NewNamespace("/v1",
		beego.NSRouter("/properties", &controllers.PropertyController{}),
		beego.NSRouter("/properties/:location", &controllers.PropertyController{}),
	)
	beego.AddNamespace(apiV1)
}

func writeJSONError(ctx *context.Context, status int, message string) {
	ctx.Output.SetStatus(status)
	_ = ctx.Output.JSON(map[string]string{"error": message}, false, false)
}
