package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["go_project/blog/controllers:LoginController"] = append(beego.GlobalControllerRouter["go_project/blog/controllers:LoginController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"get"},
			nil})

}
