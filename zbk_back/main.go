package main

import (
	"go_project/zbk_back/controllers"
	"go_project/zbk_back/models"
	_ "go_project/zbk_back/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
}
func main() {
	//models.Test_managent_api_list_files()
	beego.Router("/home", &controllers.HomeController{})
	//自动注册路由
	beego.AutoRouter(&controllers.VideoController{})

	beego.Run()
}
