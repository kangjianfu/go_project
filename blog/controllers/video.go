package controllers

import (
	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

func (this *VideoController) Show() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "551730672@gmail.com"
	this.TplName = "video.html"

	//c.TplName = "home.html"
}
