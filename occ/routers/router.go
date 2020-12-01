package routers

import (
	"occ/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/", &controllers.MainController{})

	beego.Router("/occ/user/add", &controllers.UserController{}, "post:Add")
	beego.Router("/occ/user/:id", &controllers.UserController{}, "delete:Delete")
	beego.Router("/occ/user/update", &controllers.UserController{}, "put:Update")
	beego.Router("/occ/user/detail", &controllers.UserController{}, "get:Detail")
	beego.Router("/occ/user/list", &controllers.UserController{}, "get:List")
}
