package {{.controllerPackageName}}

import (
	"core/models"
	localMod "{{.modelSrc}}/models"
	"core/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type {{.upModelName}}Controller struct {
	beego.Controller
}


func (c *{{.upModelName}}Controller) GetAll(){
	mods, err := models.SelectAll(&localMod.{{.upModelName}}{})

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
	}else {
		c.Data["json"] = mods
	}

	c.ServeJSON()
}

func (c *{{.upModelName}}Controller) Get{{.upModelName}}(){
	param := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(param, 10, 64)

	mod := models.SelectById(&localMod.{{.upModelName}}{Id: id})

	if err != nil ||  mod == nil {
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
		c.ServeJSON()
		return
	}

	c.Data["json"] = mod
	c.ServeJSON()
}

func (c *{{.upModelName}}Controller) Insert{{.upModelName}}(){
	jsonStr := c.Input().Get("{{.tableName}}Array")

	var mapResult map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, "{{.tableName}} insert fault")
	}

    mod := localMod.New{{.upModelName}}(mapResult)
    m := models.Insert(mod)

    if &m != nil {
        c.Data["json"] = m
    }else {
        c.Data["json"] = utils.GenerateRequest(500, "{{.tableName}} insert fault")
    }

	c.ServeJSON()
}

func (c *{{.upModelName}}Controller) Delete{{.upModelName}}(){
	param := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		c.Data["json"] = utils.GenerateRequest(400, "Illegal id")
		c.ServeJSON()
		return
	}

	if models.SelectById(&localMod.{{.upModelName}}{Id: id}) == nil {
		c.Data["json"] = utils.GenerateRequest(200,  "{{.tableName}} has deleted")
		c.ServeJSON()
		return
	}

	err_u := models.Delete(&localMod.{{.upModelName}}{Id: id})

	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err_u.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "{{.tableName}} delete successful")
	}

	c.ServeJSON()
}

func (c *{{.upModelName}}Controller) DeleteMore(){
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
		mod[i] = &localMod.{{.upModelName}}{Id: idResult[i]}
	}

	err_u := models.DeleteMore(mod)

	if err_u != nil {
		c.Data["json"] = utils.GenerateRequest(500, err.Error())
	}else {
		c.Data["json"] = utils.GenerateRequest(200, "{{.tableName}} delete successful")
	}

	c.ServeJSON()
}

func (c *{{.upModelName}}Controller) InsertMore(){
	jsonStr := c.Input().Get("{{.tableName}}Array")

	var mapResult []map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err == nil {
		mods := make([]models.Model, len(mapResult))

		for i := 0; i < len(mapResult); i++ {
			mods[i] = localMod.New{{.upModelName}}(mapResult[i])
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

func (c *{{.upModelName}}Controller) Update{{.upModelName}}(){
	jsonStr := c.Input().Get("{{.tableName}}")

	var mapResult map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err == nil {
		mod := localMod.New{{.upModelName}}(mapResult)

		var e error
		if e = models.Update(mod); e == nil {
			c.Data["json"] = utils.GenerateRequest(200, "update {{.tableName}} successful")
		}else {
			c.Data["json"] = utils.GenerateRequest(500, e.Error())
		}
	}else{
		c.Data["json"] = utils.GenerateRequest(400, err.Error())
	}

	c.ServeJSON()
}

func (c *{{.upModelName}}Controller) UpdateMore(){
	jsonStr := c.Input().Get("{{.tableName}}Array")

	var mapResult []map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err == nil {
		mods := make([]models.Model, len(mapResult))

		for i := 0; i < len(mapResult); i++ {
			mods[i] = localMod.New{{.upModelName}}(mapResult[i])
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

