package routers

import (
	"core/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/*)",beego.BeforeRouter, func(ctx *context.Context) {
		if ctx.Request.URL.String() != "/generator" {
			ctx.Redirect(401, "http://127.0.0.1:8090/generator")
		}

	})

	beego.Router("/generator", &controllers.MainController{},"get:Get")
}
