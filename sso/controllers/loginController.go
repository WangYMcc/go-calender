package controllers

import (
	"core/utils"
	"github.com/astaxie/beego"
	"sso/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login(){
	username := c.Input().Get("username")
	password := c.Input().Get("password")

	sessionId := c.Ctx.GetCookie("beegosessionID")
	beego.Info(sessionId)
	if u := models.Login(username, password); u != nil {

		if token, e := utils.CreateToken(u.String()); e == nil{
			beego.Info(token)

			c.Data["json"] = utils.GenerateRequest(200, "","login successful")
		}else {
			c.Data["json"] = utils.GenerateRequest(400, e.Error(), "")
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(400, "username or password fault", "")
	}


	c.ServeJSON()
}

func (c *LoginController) Get(){
	c.Data["json"] = utils.GenerateRequest(200, "", "")
	c.ServeJSON()
}

func (c *LoginController) IsLogin(){
	token := c.Ctx.GetCookie("token")
	beego.Info(token)

	if &token == nil || token == "" {
		c.Data["json"] = utils.GenerateRequest(401, "login expired", "")
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "", token)
	}

	c.ServeJSON()
}