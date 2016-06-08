package controllers

import (
	"go_project/zbk_back/models"

	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

func (this *VideoController) Page() {
	this.TplName = "video-list.html"
}
func (this *VideoController) List() {
	resultData := models.Managent_api_list_files()
	vide_datagrid := models.Zby_model_to_datagrid_video(resultData)
	this.Data["json"] = vide_datagrid
	this.ServeJSON()
}
