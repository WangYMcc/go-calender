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
var (
	isOrmInit bool
)

func init(){
	dataSource := string(beego.AppConfig.String("mysql.user") + ":" + beego.AppConfig.String("mysql.password") + "@tcp(" + beego.AppConfig.String("mysql.host") + ":" +
		beego.AppConfig.String("mysql.port") + ")/" + beego.AppConfig.String("mysql.dbname") + "?charset=utf8");
	orm.RegisterDataBase("default", "mysql", dataSource, 30)
	orm.Debug = true
	isOrmInit = true
	beego.Debug("init")
}

func RunSyncDb(){
	orm.RunSyncdb("default", false, true)
}

func IsOrmInit() bool {
	return isOrmInit
}



