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

//微信配置信息的datagrid
type Result_json struct {
	Ret    int    `json:"ret"`
	Msg    string `json:"msg"`
	Status bool   `json:"status"`
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

//微信配置信息页面
func (this *VideoController) Wx_infos_page() {
	this.TplName = "wx-infos-list.html"

}

//微信配置信息列表
func (this *VideoController) Wx_infos_list() {
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
	this.Data["json"] = models.DataGrid_Wx_infos_list(page, rows, nil)

	this.ServeJSON()
}

//提交微信配置信息
func (this *VideoController) Add_wx_info() {
	wx_info := new(models.Wei_xin_config_info)
	result := new(Result_json)
	if err := this.ParseForm(wx_info); err != nil {
		result.Ret = 1
		result.Status = false
		result.Msg = "数据解析出错，请检查数据信息。"

	} else {

		id := models.Save_wx_info(wx_info)
		if id > 0 {
			result.Ret = 0
			result.Status = true
			result.Msg = "保存成功。"
		} else {
			result.Ret = 1
			result.Status = false
			result.Msg = "网络异常，保存失败。"
		}

	}
	log.Print("插入微信配置信息成功。")
	this.Data["json"] = result
	this.ServeJSON()
}

//所有视频信息
func (this *VideoController) All_page() {
	this.TplName = "videos-all.html"
}

//视频信息列表
func (this *VideoController) All_list_videos() {
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
	query := new(models.Query_model)
	query.Page = page
	query.Rows = rows
	this.Data["json"] = models.Get_all_videos(query)

	this.ServeJSON()
}
