package routers

import (
	"core/controllers"
	"github.com/astaxie/beego"
	"core/models"
)

func init() {
	models.GetDbOrm()
    beego.Router("/", &controllers.MainController{})
	models.GetDbOrm()
}
