package controllers

import (
	"github.com/astaxie/beego"
)

//主页的控制器
type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["title"] = "威尔斯"
	this.Data["isIndex"] = true
	this.TplName = "webpage/index.html"

}

//团队结束的控制器
type TeamController struct {
	beego.Controller
}

func (this *TeamController) Get() {
	this.Data["isTeam"] = true
	this.Data["title"] = "团队介绍"
	this.TplName = "webpage/team.html"
}
