package controllers

import (
	"core/models"
	localMod "occ/models"
	"core/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type RoleController struct {
	beego.Controller
}


func (c *RoleController) GetAll(){
	mods, err := models.SelectAll(&localMod.Role{})

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
	}else {
		c.Data["json"] = mods
	}

	c.ServeJSON()
}

func (c *RoleController) GetRole(){
	param := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(param, 10, 64)

	mod := models.SelectById(&localMod.Role{Id: id})

	if err != nil ||  mod == nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
		c.ServeJSON()
		return
	}

	c.Data["json"] = mod
	c.ServeJSON()
}

func (c *RoleController) InsertRole(){
	jsonStr := c.Input().Get("roleArray")

	var mapResult map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, "role insert fault")
	}

    mod := localMod.NewRole(mapResult)
    m := models.Insert(mod)

    if &m != nil {
        c.Data["json"] = m
    }else {
        c.Data["json"] = utils.GenerateRequest(500, "role insert fault")
    }

	c.ServeJSON()
}

func (c *RoleController) DeleteRole(){
	param := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, "Illegal id")
		c.ServeJSON()
		return
	}

	if models.SelectById(&localMod.Role{Id: id}) == nil {
		c.Data["json"] = utils.GenerateRequest(200,  "role has deleted")
		c.ServeJSON()
		return
	}

	err_u := models.Delete(&localMod.Role{Id: id})

	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err_u.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "role delete successful")
	}

	c.ServeJSON()
}

func (c *RoleController) DeleteMore(){
	param := c.Input().Get("idArray")

	var idResult []int64
	err := json.Unmarshal([]byte(param), &idResult)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
		c.ServeJSON()
		return
	}

	mod := make([]models.Model, len(idResult))

	for i := 0; i < len(mod); i++{
		mod[i] = &localMod.Role{Id: idResult[i]}
	}

	err_u := models.DeleteMore(mod)

	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "role delete successful")
	}

	c.ServeJSON()
}

func (c *RoleController) InsertMore(){
	jsonStr := c.Input().Get("roleArray")

	var mapResult []map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err == nil {
		mods := make([]models.Model, len(mapResult))

		for i := 0; i < len(mapResult); i++ {
			mods[i] = localMod.NewRole(mapResult[i])
		}

		objs, err_u := models.InsertMore(mods)

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

func (c *RoleController) UpdateRole(){
	jsonStr := c.Input().Get("role")

	var mapResult map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err == nil {
		mod := localMod.NewRole(mapResult)

		var e error
		if e = models.Update(mod); e == nil {
			c.Data["json"] = utils.GenerateRequest(200, "update role successful")
		}else {
			c.Data["json"] = utils.GenerateRequest(500, e.Error())
		}
	}else{
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
	}

	c.ServeJSON()
}

func (c *RoleController) UpdateMore(){
	jsonStr := c.Input().Get("roleArray")

	var mapResult []map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err == nil {
		mods := make([]models.Model, len(mapResult))

		for i := 0; i < len(mapResult); i++ {
			mods[i] = localMod.NewRole(mapResult[i])
		}

		objs, err_u := models.UpdateMore(mods)

		if len(objs) != 0 {
			c.Data["json"] = objs
		}else{
			c.Data["json"] = utils.GenerateRequest(500, err_u.Error())
		}
	}else{
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
	}

	c.ServeJSON()
}

