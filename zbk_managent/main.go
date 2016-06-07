package main

import (
	//"fmt"
	"go_project/zbk_managent/models"
	_ "go_project/zbk_managent/routers"

	"github.com/astaxie/beego"
)

func main() {
	//task_list := models.Get_task_list()
	models.Test_management_user_list()
	//fmt.Println(task_list)
	beego.Run()
}
