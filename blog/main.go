package main

import (
	"go_project/blog/controllers"
	"go_project/blog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()

}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	//注册前端页面的路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/index.html", &controllers.HomeController{})
	beego.Router("/team.html", &controllers.TeamController{})

	beego.AutoRouter(&controllers.LoginController{})
	//beego.Router("/login", &controllers.LoginController{})
	//自动注册路由
	beego.AutoRouter(&controllers.VideoController{})
	beego.Run()
}
