package controllers

import (
	"core/models"
	"github.com/astaxie/beego"
)

func Generate(){
	user := models.SelectAll()
	beego.Debug(user)
}