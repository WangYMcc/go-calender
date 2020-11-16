package sysInit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*type Type int
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
}*/

func init(){
	confMap := GetConf()
	dataSource := string(confMap["user"] + ":" + confMap["password"]+ "@tcp(" + confMap["host"] + ":" + confMap["port"] + ")/" + confMap["dbname"] + "?charset=utf8");
	orm.RegisterDataBase("default", "mysql", dataSource, 30)
	orm.Debug = true

	orm.RunSyncdb("default", false, true)

	beego.Debug("init")
}

func RunSyncDb(){
	orm.RunSyncdb("default", false, true)
}



