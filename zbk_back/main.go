package main

import (
	"go_project/zbk_back/controllers"
	_ "go_project/zbk_back/routers"

	"github.com/astaxie/beego"
)

func main() {
	//models.Test_management_user_list()
	beego.Router("/home", &controllers.HomeController{})
	//自动注册路由
	beego.AutoRouter(&controllers.VideoController{})
	beego.Run()
}
