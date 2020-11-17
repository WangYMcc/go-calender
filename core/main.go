package main

import (
	"core/models"
	_ "core/models"
	_ "core/sysInit"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

func main() {
	//controllers.Generate()
	go func() {
		time.Sleep(3 * time.Second)
		beego.Debug("start")
		user := make([]models.User, 1000)

		for i := 0; i < len(user); i++ {
			user[i] = models.User{Username: "xm" + fmt.Sprint(i), Password: fmt.Sprint(i,i,i,i,i,i)}
		}

		models.InsertMore(user)
		beego.Debug("ok")
	}()
	beego.Run()
}

