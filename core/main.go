package main

import (
	_ "core/models"
	_ "core/routers"
	_ "core/sysInit"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	beego.Run()
}

