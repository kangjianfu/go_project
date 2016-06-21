package controllers

import (
	"go_project/zbk_back/models"
	"log"
	"strconv"

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

//举报页面
func (this *VideoController) Report_page() {
	this.TplName = "report-list.html"
}

//举报列表
func (this *VideoController) Report_list() {
	page, err := strconv.Atoi(this.Input().Get("page"))
	if err != nil {
		log.Print("页码不对")
		return
	}
	rows, err := strconv.Atoi(this.Input().Get("rows"))
	if err != nil {
		log.Print("每页条数不对")
		return
	}
	this.Data["json"] = models.Datagrid_report_info(page, rows)
	this.ServeJSON()

}

//服务码页面
func (this *VideoController) Code_page() {
	this.TplName = "codes-list.html"

}

//服务码列表
func (this *VideoController) Code_list() {
	page, err := strconv.Atoi(this.Input().Get("page"))
	if err != nil {
		log.Print("页码不对")
		return
	}
	rows, err := strconv.Atoi(this.Input().Get("rows"))
	if err != nil {
		log.Print("每页条数不对")
		return
	}
	this.Data["json"] = models.DataGrid_service_codes_list(page, rows, nil)

	this.ServeJSON()
}
