package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//实现 controller的登录方法
func (this *LoginController) Get() {
	//登录页面不设置
	//this.Data["xsrfdata"] = this.XSRFFormHTML()
	this.Data["title"] = "登录页面"

	//跳转登录页面
	this.TplName = "login.html"

}
