package main

import (
	"bytes"
	"github.com/astaxie/beego"
	"html/template"
	"os"
)

func main() {
	temp := template.Must(template.ParseFiles("views/user.tpl"))
	f, _ := os.OpenFile("views/test.tpl", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	defer f.Close()
	w := bytes.NewBuffer(nil)

	beego.Info(temp.Execute(w, map[string]interface{}{
		"packageName": "models",
		"upModelName": "User",
		"tableName": "user",
		"paramMap": map[string]map[string]interface{}{
			"Name": {
				"type":"string",
				"prop":[]string{"column(name)"}},
			"Password": {
				"type":"string",
				"prop":[]string{"column(password)"}},
			"Account": {
				"type":"string",
				"prop":[]string{"column(account)", ""}},
			},
		}))

	beego.Info(w.String())
	f.Write(w.Bytes())
}
