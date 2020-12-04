package controllers

import (
	"core/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	localMod "occ/models"
	coreMod "core/models"
	"strconv"
)

type RoleController struct {
	beego.Controller
}

func (c *RoleController)Add(){
	var m localMod.Role

	if e := json.Unmarshal(c.Ctx.Input.RequestBody, &m); e == nil {
		obj, err := coreMod.Insert(&m)

		if err == nil {
			c.Data["json"] = utils.GenerateRequestWithObj(utils.HTTP_OK, "role insert successful", obj)
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, err.Error())
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, e.Error())
	}

	c.ServeJSON()
}

func (c *RoleController) Delete(){
	param := c.Input().Get("id")
	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, "Illegal id")
		c.ServeJSON()
		return
	}

	if coreMod.SelectById(&localMod.Role{Id: id}) == nil {
		c.Data["json"] = utils.GenerateRequest(200,  "role does not exist")
		c.ServeJSON()
		return
	}

	err_u := coreMod.Delete(&localMod.Role{Id: id})
	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err_u.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "user delete successful")
	}

	c.ServeJSON()
}

func (c *RoleController) DeleteBatch(){
	param := c.Input().Get("idArray")

	var idResult []int64
	err := json.Unmarshal([]byte(param), &idResult)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
		c.ServeJSON()
		return
	}

	role := make([]coreMod.Model, len(idResult))

	for i := 0; i < len(role); i++{
		role[i] = &localMod.Role{Id: idResult[i]}
	}

	err_u := coreMod.DeleteMore(role)

	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "user delete successful")
	}

	c.ServeJSON()
}

func (c *RoleController)Update(){
	var m localMod.Role

	if e := json.Unmarshal(c.Ctx.Input.RequestBody, &m); e == nil {
		err := coreMod.Update(&m)

		if err == nil {
			c.Data["json"] = utils.GenerateRequestWithObj(utils.HTTP_OK, "role insert successful", m)
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, err.Error())
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, e.Error())
	}

	c.ServeJSON()
}

func (c *RoleController)List(){
	page := utils.ChangeValInt(c.Input().Get("page"))
	size := utils.ChangeValInt(c.Input().Get("size"))

	if page != -1 && size != -1 {
		role, err := coreMod.SelectAll(&localMod.Role{})

		if err == nil {
			c.Data["json"] = utils.GenerateRequestWithList(utils.HTTP_OK, "", utils.StartPage(page, size, role))
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, err.Error())
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, "page or size Illegal")
	}

	c.ServeJSON()
}

func (c *RoleController)Detail(){
	id := utils.ChangeValInt64(c.Input().Get("id"))

	if id != -1 {
		role := coreMod.SelectById(&localMod.Role{Id:id})

		if role != nil {
			c.Data["json"] = utils.GenerateRequestWithObj(utils.HTTP_OK, "", role)
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_NOT_FOUND, "could not found the role")
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, "id Illegal")
	}

	c.ServeJSON()
}

func (c *RoleController)ListByKey(){
	page := utils.ChangeValInt(c.Input().Get("page"))
	size := utils.ChangeValInt(c.Input().Get("size"))
	key := c.Input().Get("key")
	val := c.Input().Get("val")

	if page != -1 && size != -1 {
		var role []localMod.Role
		err := coreMod.SelectAllByKey(&localMod.Role{}, key, val, &role)

		if err == nil {
			mods := make([]coreMod.Model, len(role))

			for i := 0; i < len(role); i++ {
				mods[i] = &role[i]
			}

			c.Data["json"] = utils.GenerateRequestWithList(utils.HTTP_OK, "", utils.StartPage(page, size, mods))
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, err.Error())
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, "page or size Illegal")
	}

	c.ServeJSON()
}