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
	//注册首页的路由
	beego.Router("/", &controllers.HomeController{})
	//注册登录的路由
	//beego.Router("/login", )
	//注解注册路由 loginController
	//beego.auto()
	beego.AutoRouter(&controllers.LoginController{})
	//beego.Router("/login", &controllers.LoginController{})
	//自动注册路由
	beego.AutoRouter(&controllers.VideoController{})
	beego.Run()
}
