package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Type int
const(
	Integer Type = iota
	Varchar
	Text
)

type Table struct {
	pirmary bool
	isNull bool
	valType int
	length int
	unsign bool
	zeroFill bool
	autoIncrement bool
}
var (
	o orm.Ormer
)

func GetDbOrm() orm.Ormer{
	if o != nil {
		return o
	}

	confMap := GetConf("mysql")
	dataSource := string(confMap["user"] + ":" + confMap["password"]+ "@tcp(" + confMap["host"] + ":" + confMap["port"] + ")/" + confMap["dbname"] + "?charset=utf8");

	orm.RegisterDataBase("default", "mysql", dataSource, 30)
	o = orm.NewOrm()

	return o
}


