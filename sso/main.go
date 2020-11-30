package main

import (
	_ "core/sysInit/sql"
	_ "core/sysInit/redis"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "sso/routers"
)

func main() {
	beego.Run()
}

