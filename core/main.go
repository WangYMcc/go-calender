package main

import (
	"core/models"
	_ "core/sysInit"
	"github.com/astaxie/beego"
)

func main() {
	user := models.SelectAll()
	beego.Debug(user)
	//controllers.Generate()
	//beego.Run()
}

