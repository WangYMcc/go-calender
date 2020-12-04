package {{.controllerPackageName}}

import (
	"core/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	localMod "{{.modelSrc}}/models"
	coreMod "core/models"
	"strconv"
)

type {{.upModelName}}Controller struct {
	beego.Controller
}

func (c *{{.upModelName}}Controller)Add(){
	var m localMod.{{.upModelName}}

	if e := json.Unmarshal(c.Ctx.Input.RequestBody, &m); e == nil {
		obj, err := coreMod.Insert(&m)

		if err == nil {
			c.Data["json"] = utils.GenerateRequestWithObj(utils.HTTP_OK, "{{.tableName}} insert successful", obj)
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, err.Error())
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, e.Error())
	}

	c.ServeJSON()
}

func (c *{{.upModelName}}Controller) Delete(){
	param := c.Input().Get("id")
	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, "Illegal id")
		c.ServeJSON()
		return
	}

	if coreMod.SelectById(&localMod.{{.upModelName}}{Id: id}) == nil {
		c.Data["json"] = utils.GenerateRequest(200,  "{{.tableName}} does not exist")
		c.ServeJSON()
		return
	}

	err_u := coreMod.Delete(&localMod.{{.upModelName}}{Id: id})
	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err_u.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "user delete successful")
	}

	c.ServeJSON()
}

func (c *{{.upModelName}}Controller) DeleteBatch(){
	param := c.Input().Get("idArray")

	var idResult []int64
	err := json.Unmarshal([]byte(param), &idResult)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
		c.ServeJSON()
		return
	}

	{{.tableName}} := make([]coreMod.Model, len(idResult))

	for i := 0; i < len({{.tableName}}); i++{
		{{.tableName}}[i] = &localMod.{{.upModelName}}{Id: idResult[i]}
	}

	err_u := coreMod.DeleteMore({{.tableName}})

	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "user delete successful")
	}

	c.ServeJSON()
}

func (c *{{.upModelName}}Controller)Update(){
	var m localMod.{{.upModelName}}

	if e := json.Unmarshal(c.Ctx.Input.RequestBody, &m); e == nil {
		err := coreMod.Update(&m)

		if err == nil {
			c.Data["json"] = utils.GenerateRequestWithObj(utils.HTTP_OK, "{{.tableName}} insert successful", m)
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, err.Error())
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, e.Error())
	}

	c.ServeJSON()
}

func (c *{{.upModelName}}Controller)List(){
	page := utils.ChangeValInt(c.Input().Get("page"))
	size := utils.ChangeValInt(c.Input().Get("size"))

	if page != -1 && size != -1 {
		{{.tableName}}, err := coreMod.SelectAll(&localMod.{{.upModelName}}{})

		if err == nil {
			c.Data["json"] = utils.GenerateRequestWithList(utils.HTTP_OK, "", utils.StartPage(page, size, {{.tableName}}))
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, err.Error())
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, "page or size Illegal")
	}

	c.ServeJSON()
}

func (c *{{.upModelName}}Controller)Detail(){
	id := utils.ChangeValInt64(c.Input().Get("id"))

	if id != -1 {
		{{.tableName}} := coreMod.SelectById(&localMod.{{.upModelName}}{Id:id})

		if {{.tableName}} != nil {
			c.Data["json"] = utils.GenerateRequestWithObj(utils.HTTP_OK, "", {{.tableName}})
		}else {
			c.Data["json"] = utils.GenerateRequest(utils.HTTP_NOT_FOUND, "could not found the {{.tableName}}")
		}
	}else {
		c.Data["json"] = utils.GenerateRequest(utils.HTTP_SERVER_ERROR, "id Illegal")
	}

	c.ServeJSON()
}

func (c *{{.upModelName}}Controller)ListByKey(){
	page := utils.ChangeValInt(c.Input().Get("page"))
	size := utils.ChangeValInt(c.Input().Get("size"))
	key := c.Input().Get("key")
	val := c.Input().Get("val")

	if page != -1 && size != -1 {
		var {{.tableName}} []localMod.{{.upModelName}}
		err := coreMod.SelectAllByKey(&localMod.{{.upModelName}}{}, key, val, &{{.tableName}})

		if err == nil {
			mods := make([]coreMod.Model, len({{.tableName}}))

			for i := 0; i < len({{.tableName}}); i++ {
				mods[i] = &{{.tableName}}[i]
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