package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_project/blog/controllers"
	"go_project/blog/models"
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
	beego.Router("/login", &controllers.LoginController{})
	//自动注册路由
	beego.AutoRouter(&controllers.VideoController{})
	beego.Run()
}
