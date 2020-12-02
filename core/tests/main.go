package main

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"os"
	"strings"
)

func main() {
	temp := template.Must(template.ParseFiles("views/route.tpl"))
	generate := map[string]interface{}{
		"upModelName": beego.AppConfig.String("modelName"),
		"tableName": beego.AppConfig.String("tableName"),
		"modelSrc": beego.AppConfig.String("modelSrc"),
	}

	f, _ := os.OpenFile(fmt.Sprint("views/route.txt"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
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
