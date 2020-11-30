package controllers

import (
	"core/utils"
	"core/utils/jwt"
	"github.com/astaxie/beego"
	"sso/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login(){
	username := c.Input().Get("username")
	password := c.Input().Get("password")

	if u := models.Login(username, password); u != nil {

		if accessToken, e := jwt.CreateAccessToken(u.String()); e == nil{
			if flushToken, e2 := jwt.CreateFlushToken(u.String()); e2 == nil {

				beego.Info(accessToken)
				beego.Info(flushToken)
				c.Ctx.SetCookie("access_token", accessToken)
				c.Ctx.SetCookie("flush_token", flushToken)
				c.Data["json"] = utils.GenerateRequest(200, "login successful")
			}else {
				c.Data["json"] = utils.GenerateRequest(400, e2.Error())
			}
		}else {
			c.Data["json"] = utils.GenerateRequest(400, e.Error())
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(400, "username or password fault")
	}


	c.ServeJSON()
}

func (c *LoginController) Get(){
	c.Data["json"] = utils.GenerateRequest(200, "SSO")
	c.ServeJSON()
}

func (c *LoginController) IsLogin(){
	accessToken := c.Ctx.GetCookie("access_token")
	flushToken := c.Ctx.GetCookie("access_token")
	beego.Info(accessToken)
	beego.Info(flushToken)

	if flushToken == "" || accessToken == "" {
		c.Data["json"] = utils.GenerateRequest(401, "login expired")
	}else {
		c.Data["json"] = utils.GenerateRequest(200, accessToken)
	}

	c.ServeJSON()
}
