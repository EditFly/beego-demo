package routers

import (
	"beegoweb/app/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.CMSController{})
}
