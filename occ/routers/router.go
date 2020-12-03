package routers

import (
	"github.com/astaxie/beego"
	"occ/controllers"
)

func init() {

    beego.Router("/", &controllers.MainController{})

	beego.Router("/occ/user/add", &controllers.UserController{}, "post:Add")
	beego.Router("/occ/user/delete", &controllers.UserController{}, "get:Delete")
	beego.Router("/occ/user/deleteBatch", &controllers.UserController{}, "get:DeleteBatch")
	beego.Router("/occ/user/update", &controllers.UserController{}, "put:Update")
	beego.Router("/occ/user/detail", &controllers.UserController{}, "get:Detail")
	beego.Router("/occ/user/list", &controllers.UserController{}, "get:List")
	/*beego.Router("/occ/user/listByKey", &controllers.UserController{}, "get:ListByKey")*/

}
