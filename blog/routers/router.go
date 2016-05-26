package routers

import (
	"github.com/astaxie/beego"
	"go_project/blog/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
}
