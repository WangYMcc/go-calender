package main

import (
	"core/controllers"
	_ "core/models"
	_ "core/sysInit"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	beego.InsertFilter("/*)",beego.BeforeRouter, func(ctx *context.Context) {
		r, _ := ctx.Request.Cookie("beegosessionID")
		if r != nil {
			beego.Info(r.Value)
		}
	})

	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/all", &controllers.UserController{}, "get:GetAll")
	beego.Router("/user/:id", &controllers.UserController{}, "get:GetUser")
	beego.Router("/user/insert", &controllers.UserController{}, "put:InsertUser")
	beego.Router("/user/insertmore", &controllers.UserController{}, "put:InsertMore")
	beego.Router("/user/:id", &controllers.UserController{}, "delete:DeleteUser")
	beego.Router("/user/deletemore", &controllers.UserController{}, "delete:DeleteMore")
	beego.Router("/user/update", &controllers.UserController{}, "post:UpdateUser")
	beego.Router("/user/updatemore", &controllers.UserController{}, "post:UpdateMore")
	beego.Run()
}

