package controllers

import (
	"core/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"sso/models"
	coreMod "core/models"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (c *UserController)Add(){
	var m models.User

	if e := json.Unmarshal(c.Ctx.Input.RequestBody, &m); e == nil {
		obj, err := coreMod.Insert(&m)

		if err == nil {
			c.Data["json"] = utils.GenerateRequestWithObj(utils.HTTP_OK, "user insert successful", obj)
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, err.Error())
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, e.Error())
	}

	c.ServeJSON()
}

func (c *UserController) Delete(){
	param := c.Input().Get("id")
	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, "Illegal id")
		c.ServeJSON()
		return
	}

	if coreMod.SelectById(&models.User{Id: id}) == nil {
		c.Data["json"] = utils.GenerateRequest(200,  "user does not exist")
		c.ServeJSON()
		return
	}

	err_u := coreMod.Delete(&models.User{Id: id})
	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err_u.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "user delete successful")
	}

	c.ServeJSON()
}

func (c *UserController) DeleteBatch(){
	param := c.Input().Get("idArray")

	var idResult []int64
	err := json.Unmarshal([]byte(param), &idResult)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
		c.ServeJSON()
		return
	}

	user := make([]coreMod.Model, len(idResult))

	for i := 0; i < len(user); i++{
		user[i] = &models.User{Id: idResult[i]}
	}

	err_u := coreMod.DeleteMore(user)

	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "user delete successful")
	}

	c.ServeJSON()
}

func (c *UserController)Update(){
	var m models.User

	if e := json.Unmarshal(c.Ctx.Input.RequestBody, &m); e == nil {
		err := coreMod.Update(&m)

		if err == nil {
			c.Data["json"] = utils.GenerateRequestWithObj(utils.HTTP_OK, "user insert successful", m)
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, err.Error())
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, e.Error())
	}

	c.ServeJSON()
}

func (c *UserController)List(){
	page := utils.ChangeValInt(c.Input().Get("page"))
	size := utils.ChangeValInt(c.Input().Get("size"))

	if page != -1 && size != -1 {
		users, err := coreMod.SelectAll(&models.User{})

		if err == nil {
			c.Data["json"] = utils.GenerateRequestWithList(utils.HTTP_OK, "", utils.StartPage(page, size, users))
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, err.Error())
		}

	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, "page or size Illegal")
	}


	c.ServeJSON()
}

func (c *UserController)Detail(){
	id := utils.ChangeValInt64(c.Input().Get("id"))

	if id != -1 {
		user := coreMod.SelectById((&models.User{Id:id}).ToModel())

		if user != nil {
			c.Data["json"] = utils.GenerateRequestWithObj(utils.HTTP_OK, "", user)
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_NOT_FOUND, "could not found the user")
		}

	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, "id Illegal")
	}


	c.ServeJSON()
}