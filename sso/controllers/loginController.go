package controllers

import (
	"core/sysInit/redis"
	"core/utils"
	"core/utils/jwt"
	"fmt"
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

				redis.SET(fmt.Sprint("loginUser:", u.Id), u)

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
	accessToken := c.Ctx.GetCookie("access_token")

	_, err := jwt.ParseToken(accessToken)
	if err != nil {
		c.Data["json"] = utils.GenerateRequest(401, "login expired")
		beego.Info("login expired")
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "SSO")
	}
	c.ServeJSON()
}

func (c *LoginController) IsLogin(){
	accessToken := c.Ctx.GetCookie("access_token")
	beego.Info(accessToken)

	if accessToken == "" {
		c.Data["json"] = utils.GenerateRequest(401, "login expired")
	}else {
		_, err := jwt.ParseToken(accessToken)
		if err != nil {
			c.Data["json"] = utils.GenerateRequest(401, "login expired")
		}else {
			c.Data["json"] = utils.GenerateRequest(200, accessToken)
		}
	}

	c.ServeJSON()
}
