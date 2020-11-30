package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"sso/controllers"
)

func init() {
	beego.InsertFilter("/*",beego.BeforeRouter, func(ctx *context.Context) {
		if ctx.Request.URL.String() == "/sso/to" {

		}

	})

	beego.Router("/sso/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/sso/isLogin", &controllers.LoginController{}, "get:IsLogin")
	beego.Router("/", &controllers.LoginController{})
	/*beego.Router("/sso/user/all", &controllers.UserController{}, "get:GetAll")
	beego.Router("/sso/user/:id", &controllers.UserController{}, "get:GetUser")
	beego.Router("/sso/user/insert", &controllers.UserController{}, "put:InsertUser")
	beego.Router("/sso/user/insertmore", &controllers.UserController{}, "put:InsertMore")
	beego.Router("/sso/user/:id", &controllers.UserController{}, "delete:DeleteUser")
	beego.Router("/sso/user/deletemore", &controllers.UserController{}, "delete:DeleteMore")
	beego.Router("/sso/user/update", &controllers.UserController{}, "post:UpdateUser")
	beego.Router("/sso/user/updatemore", &controllers.UserController{}, "post:UpdateMore")*/
}
