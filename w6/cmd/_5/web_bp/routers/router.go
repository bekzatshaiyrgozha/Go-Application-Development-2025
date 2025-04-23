package routers

import (
	"web_bp/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/custom", &controllers.MainController{}, "*:Custom")
}
