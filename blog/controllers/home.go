package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "551730672@gmail.com"
	c.TplName = "index.html"
}
