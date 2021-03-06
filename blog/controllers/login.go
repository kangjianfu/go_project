package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (a *LoginController) Prepare() {
	a.EnableXSRF = false
}

//@router /login页面
func (this *LoginController) In() {
	//登录页面不设置
	this.Data["xsrfdata"] = this.XSRFFormHTML()
	this.Data["title"] = "登录页面"
	//跳转登录页面

	this.TplName = "login.html"

}

//登录操作
func (this *LoginController) Login() {
	user_name := this.Input().Get("user_name")
	password := this.Input().Get("pwd")
	//email := this.Input().Get("email")
	//查询数据
	//设置cookie
	maxAge := 0
	this.Ctx.SetCookie("user_name", user_name, maxAge, "/")
	this.Ctx.SetCookie("user_name", password, maxAge, "/")
}

//注册页面
func (this *LoginController) Zc() {
	this.Data["title"] = "注册"
	this.TplName = "regist.html"
}
