package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"property_listing_api/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/properties", &controllers.PropertyController{})
}
