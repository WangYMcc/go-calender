package main

import (
	_ "core/sysInit"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "sso/routers"
)

func main() {
	beego.Run()
}

