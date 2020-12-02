package generator

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"os"
	"strings"
)

func GenerateModel(){
	temp := template.Must(template.ParseFiles("views/model.tpl"))
	params := strings.Split(beego.AppConfig.String("keyParam"), "|")
	paramMap := make(map[string]map[string]interface{})

	for i := 0; i < len(params); i++{
		prop := strings.Split(params[i], ";")

		paramMap[prop[0]] = make(map[string]interface{})
		if prop[1] == "" || prop[1] == "string" {
			paramMap[prop[0]]["type"] = "string"
			paramMap[prop[0]]["int"] = false
		}else {
			paramMap[prop[0]]["type"] = prop[1]
			paramMap[prop[0]]["int"] = true
		}

		paramMap[prop[0]]["low"] = strings.ToLower(prop[0])

		if prop[2] == "" {
			prop[2] += fmt.Sprint("column(", strings.ToLower(prop[0]), ")")
		}else {
			pro := fmt.Sprint("column(", strings.ToLower(prop[0]), "),", prop[2])
			prop[2] = pro
		}

		paramMap[prop[0]]["prop"] = strings.Split(prop[2], ",")
	}

	generate := map[string]interface{}{
		"modelPackageName": beego.AppConfig.String("modelPackageName"),
		"upModelName": beego.AppConfig.String("modelName"),
		"tableName": beego.AppConfig.String("tableName"),
		"paramMap": paramMap,
	}
	f, _ := os.OpenFile(fmt.Sprint("../", beego.AppConfig.String("modelSrc"), "/",beego.AppConfig.String("modelPackageName"), "/",
		beego.AppConfig.String("tableName"),".go"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	w := bytes.NewBuffer(nil)

	temp.Execute(w, generate)
	/*beego.Info(temp.Execute(w, map[string]interface{}{
	"packageName": "models",
	"upModelName": "Pser",
	"tableName": "pser",
	"paramMap": map[string]map[string]interface{}{
		"Name": {
			"type":"string",
			"low": "name",
			"prop":[]string{"column(name)"}},
		"Password": {
			"type":"string",
			"low": "password",
			"prop":[]string{"column(password)"}},
		"Account": {
			"type":"string",
			"low": "account",
			"prop":[]string{"column(account)", "unique"}},
		},
	}))*/
	s := strings.ReplaceAll(w.String(), "&lt;", "<")
	beego.Info(s)

	f.WriteString(s)
}

func GenerateController(){
	temp := template.Must(template.ParseFiles("views/controller.tpl"))
	generate := map[string]interface{}{
		"controllerPackageName": beego.AppConfig.String("controllerPackageName"),
		"upModelName": beego.AppConfig.String("modelName"),
		"tableName": beego.AppConfig.String("tableName"),
		"modelSrc": beego.AppConfig.String("modelSrc"),
	}

	f, _ := os.OpenFile(fmt.Sprint("../", beego.AppConfig.String("modelSrc"), "/",beego.AppConfig.String("controllerPackageName"), "/",
		beego.AppConfig.String("tableName"),"Controller.go"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	w := bytes.NewBuffer(nil)

	temp.Execute(w, generate)
	/*beego.Info(temp.Execute(w, map[string]interface{}{
	"packageName": "models",
	"upModelName": "Pser",
	"tableName": "pser",
	"paramMap": map[string]map[string]interface{}{
		"Name": {
			"type":"string",
			"low": "name",
			"prop":[]string{"column(name)"}},
		"Password": {
			"type":"string",
			"low": "password",
			"prop":[]string{"column(password)"}},
		"Account": {
			"type":"string",
			"low": "account",
			"prop":[]string{"column(account)", "unique"}},
		},
	}))*/
	s := strings.ReplaceAll(w.String(), "&lt;", "<")
	beego.Info(s)

	f.WriteString(s)
}