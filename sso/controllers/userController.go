package controllers

import (
	"core/models"
	"core/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	models2 "sso/models"
	"strconv"
)

type UserController struct {
	beego.Controller
}


func (c *UserController) GetAll(){
	users, err := models.SelectAll(&models2.User{})

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
	}else {
		c.Data["json"] = users
	}

	c.ServeJSON()

}

func (c *UserController) GetUser(){
	param := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(param, 10, 64)

	user := models.SelectById(&models2.User{Id: id})

	if err != nil ||  user == nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
		c.ServeJSON()
		return
	}

	c.Data["json"] = user
	c.ServeJSON()
}

func (c *UserController) InsertUser(){
	jsonStr := c.Input().Get("userArray")

	var mapResult map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, "user insert fault")
	}

	u, err := models.AccurateQuery("username", mapResult["username"], &models2.User{});

	if  u == nil && err == nil {
		user := models2.NewUser(mapResult)
		u, _ := models.Insert(user)

		if &u != nil {
			c.Data["json"] = u
		}else {
			c.Data["json"] = utils.GenerateRequest(500, "user insert fault")
		}

	}else {
		c.Data["json"] = utils.GenerateRequest(400, "username is repeated")
	}

	c.ServeJSON()
}

func (c *UserController) DeleteUser(){
	param := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, "Illegal id")
		c.ServeJSON()
		return
	}

	if models.SelectById(&models2.User{Id: id}) == nil {
		c.Data["json"] = utils.GenerateRequest(200,  "user has deleted")
		c.ServeJSON()
		return
	}

	err_u := models.Delete(&models2.User{Id: id})

	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err_u.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "user delete successful")
	}

	c.ServeJSON()
}

func (c *UserController) DeleteMore(){
	param := c.Input().Get("idArray")

	var idResult []int64
	err := json.Unmarshal([]byte(param), &idResult)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
		c.ServeJSON()
		return
	}

	user := make([]models.Model, len(idResult))

	for i := 0; i < len(user); i++{
		user[i] = &models2.User{Id: idResult[i]}
	}

	err_u := models.DeleteMore(user)

	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "user delete successful")
	}

	c.ServeJSON()
}

func (c *UserController) InsertMore(){
	jsonStr := c.Input().Get("userArray")

	var mapResult []map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err == nil {
		users := make([]models.Model, len(mapResult))

		for i := 0; i < len(mapResult); i++ {
			users[i] = models2.NewUser(mapResult[i])
		}

		objs, err_u := models.InsertMore(users)

		if objs != nil {
			c.Data["json"] = objs
		}else{
			c.Data["json"] = utils.GenerateRequest(500, err_u.Error())
		}
	}else{
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
	}

	c.ServeJSON()
}

func (c *UserController) UpdateUser(){
	jsonStr := c.Input().Get("user")

	var mapResult map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err == nil {
		user := models2.NewUser(mapResult)

		var e error
		if e = models.Update(user); e == nil {
			c.Data["json"] = utils.GenerateRequest(200, "update user successful")
		}else {
			c.Data["json"] = utils.GenerateRequest(500, e.Error())
		}
	}else{
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
	}

	c.ServeJSON()
}

func (c *UserController) UpdateMore(){
	jsonStr := c.Input().Get("userArray")

	var mapResult []map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err == nil {
		users := make([]models.Model, len(mapResult))

		for i := 0; i < len(mapResult); i++ {
			users[i] = models2.NewUser(mapResult[i])
		}

		users, err_u := models.UpdateMore(users)

		if len(users) != 0 {
			c.Data["json"] = users
		}else{
			c.Data["json"] = utils.GenerateRequest(500, err_u.Error())
		}
	}else{
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
	}

	c.ServeJSON()
}