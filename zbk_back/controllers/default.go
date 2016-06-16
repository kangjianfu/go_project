package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "沃安科技"
	c.Data["Email"] = "沃安科技@qq.com"
	c.TplName = "video_demo.html"
}
