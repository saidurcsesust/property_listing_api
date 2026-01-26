package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"property_listing_api/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	apiV1 := beego.NewNamespace("/v1",
		beego.NSRouter("/properties/:location", &controllers.PropertyController{}),
	)
	beego.AddNamespace(apiV1)
}
