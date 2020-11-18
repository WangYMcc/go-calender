package controllers

import (
	"core/models"
	"core/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type UserController struct {
	beego.Controller
}


func (c *UserController) GetAll(){
	user := models.SelectAll()

	c.Data["json"] = user
	c.ServeJSON()

}

func (c *UserController) GetUser(){
	param := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(param, 10, 64)

	user := models.SelectById(id)

	if err != nil ||  user == nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
		c.ServeJSON()
		return
	}

	c.Data["json"] = user
	c.ServeJSON()
}

func (c *UserController) InsertUser(){
	username := c.Input().Get("username")
	password := c.Input().Get("password")

	if users,err := models.Query("username", username); len(users) == 0 && err == nil {
		user := models.User{Username: username, Password: password}
		user = *user.Insert()

		if &user != nil {
			c.Data["json"] = user
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

	if models.SelectById(id) == nil {
		c.Data["json"] = utils.GenerateRequest(200, "user has deleted")
		c.ServeJSON()
		return
	}

	err_u := models.User{Id: id}.Delete()

	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err_u.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "")
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

	user := make([]models.User, len(idResult))

	for i := 0; i < len(user); i++{
		user[i] = models.User{Id: idResult[i]}
	}

	err_u := models.DeleteMore(user)

	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err_u.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "")
	}

	c.ServeJSON()
}

func (c *UserController) InsertMore(){
	jsonStr := c.Input().Get("userArray")

	var mapResult []map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err == nil {
		users := make([]models.User, len(mapResult))

		for i := 0; i < len(mapResult); i++ {
			users[i] = models.User{Username: mapResult[i]["username"], Password: mapResult[i]["password"]}
		}

		users, err_u := models.InsertMore(users)

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

func (c *UserController) UpdateUser(){
	jsonStr := c.Input().Get("user")

	var mapResult map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err == nil {
		id, err_ := strconv.ParseInt(mapResult["id"], 10, 64)

		if err_ == nil {
			users := models.User{Id:id, Username: mapResult["username"], Password: mapResult["password"]}

			var e error
			if e = users.Update(); e == nil {
				c.Data["json"] = utils.GenerateRequest(200, "")
			}else {
				c.Data["json"] = utils.GenerateRequest(500, e.Error())
			}

		}else {
			c.Data["json"] = utils.GenerateRequest(500, err_.Error())
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
		users := make([]models.User, len(mapResult))

		for i := 0; i < len(mapResult); i++ {
			id, _ := strconv.ParseInt(mapResult[i]["id"], 10, 64)
			users[i] = models.User{Id: id, Username: mapResult[i]["username"], Password: mapResult[i]["password"]}
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