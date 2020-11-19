package routers

import (
	"core/controllers"
	"github.com/astaxie/beego"
)

func init() {
	/*beego.InsertFilter("/)",beego.BeforeRouter, func(ctx *context.Context) {
		//ctx.Redirect(401, "/")
	})*/

	beego.Router("/", &controllers.MainController{},"get:Get")
}
