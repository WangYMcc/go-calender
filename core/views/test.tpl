package models

import(
    "core/models"
    "core/sysInit/redis"
    "core/sysInit/sql"
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "strconv"
    "sync"
    "time"
)

var (
	locl User
)

type User} struct {
    mu sync.Mutex
    Id  int64 `orm:"column(id);unique;pk"`
    Name  string  `orm:"column(name);"`
    Password  string  `orm:"column(password);"`
}