package main

import (
	"github.com/astaxie/beego"
)

type HomeContorller struct {
	beego.Controller
}

func (this *HomeContorller) Get() {
	this.Ctx.WriteString("hellow 你好 世界")
}

func main() {
	//注册路由
	beego.Router("/", &HomeContorller{})
	//开始运行
	beego.Run()
}
