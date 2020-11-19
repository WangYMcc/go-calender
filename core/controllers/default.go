package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	sessionId := c.Ctx.GetCookie("beegosessionID")

	req := httplib.Get("http://127.0.0.1:8000/sso/isLogin")
	str, err := req.String()

	if err != nil {
		beego.Error(err)
	}
	beego.Info(sessionId)
	beego.Info(str)
	c.Data["info"] = "Welcome to use core"
	c.ServeJSON()
}
