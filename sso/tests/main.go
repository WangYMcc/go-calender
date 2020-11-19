package main

import (
	"github.com/astaxie/beego"
	"sso/models"
)


func main(){
	beego.Info(models.User{Id: 1, Username: "1", UserKey: "1", Password: "1"})
}

