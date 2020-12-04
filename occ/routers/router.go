package routers

import (
	"github.com/astaxie/beego"
	"occ/controllers"
)

func init() {

    beego.Router("/", &controllers.MainController{})

	/*beego.Router("/occ/user/add", &controllers.UserController{}, "post:Add")
	beego.Router("/occ/user/delete", &controllers.UserController{}, "get:Delete")
	beego.Router("/occ/user/deleteBatch", &controllers.UserController{}, "get:DeleteBatch")
	beego.Router("/occ/user/update", &controllers.UserController{}, "put:Update")
	beego.Router("/occ/user/detail", &controllers.UserController{}, "get:Detail")
	beego.Router("/occ/user/list", &controllers.UserController{}, "get:List")
	beego.Router("/occ/user/listByKey", &controllers.UserController{}, "get:ListByKey")*/

	beego.Router("/occ/role/add", &controllers.RoleController{}, "post:Add")
	beego.Router("/occ/role/delete", &controllers.RoleController{}, "get:Delete")
	beego.Router("/occ/role/deleteBatch", &controllers.RoleController{}, "get:DeleteBatch")
	beego.Router("/occ/role/update", &controllers.RoleController{}, "put:Update")
	beego.Router("/occ/role/detail", &controllers.RoleController{}, "get:Detail")
	beego.Router("/occ/role/list", &controllers.RoleController{}, "get:List")
	beego.Router("/occ/role/listByKey", &controllers.RoleController{}, "get:ListByKey")

}
