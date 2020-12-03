package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "occ/routers"
)

func main() {
	beego.Run()
}

