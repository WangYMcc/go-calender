package models

import (
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

//定义数据库表结构体
type User struct {
	mu sync.Mutex
	Id       int64 `orm:"column(id);unique;pk"`
	UserKey string `orm:"column(userkey);unique;"`
	Username string `orm:"column(username);unique;"`
	Password string `orm:"column(password);null"`
}

func (u *User)Obj() interface{} {
	return u
}

func (u *User)SetId(id int64){
	u.Id = id
}

func (u *User)GetLock() *sync.Mutex{
	return &u.mu
}

//设置表名
func (u *User) TableName() string {
	return "user"
}

func (u *User)GetId() int64{
	return u.Id
}

func (u *User)ToModel() models.Model {
	return u
}


func (u *User)QueryKey() string{
	return "id, userkey, username, password"
}

func (u *User)QueryResult(maps []orm.Params) ([]models.Model, error)  {
	if len(maps) == 0 {
		return nil, nil
	}

	user := make([]User, len(maps))
	models := make([]models.Model, len(maps))

	for i := 0; i < len(maps); i++ {
		id, err := strconv.ParseInt(fmt.Sprint(maps[i]["id"]), 10, 64)
		user[i] = User{Id: id, Username: fmt.Sprint(maps[i]["username"]), Password: fmt.Sprint(maps[i]["password"])}
		models[i] = user[i].ToModel()

		if err != nil {
			return nil, err
		}

	}

	return models, nil
}

//实现String函数
func (u User) String() string{
	return fmt.Sprint(`{"id": `, u.Id, `, "userkey": "`, u.UserKey, `", "username": "`, u.Username, `"}`)
}

func NewUser(m map[string]string) *User {
	user := User{}

	id, err := strconv.ParseInt(m["id"], 10, 64)

	if err == nil {
		user.Id = id
	}

	user.UserKey = m["userkey"]
	user.Username = m["username"]
	user.Password = m["password"]

	return &user
}

//初始化-建表-将数据存到缓存中
func init(){
	orm.RegisterModel(new(User))
	locl = User{Id: 0}

	go func(){
		for !sql.IsOrmInit() {
			time.Sleep(1 * time.Second)
		}

		sql.RunSyncDb()

		for !redis.IsRedisCacheInit() {
			time.Sleep(1 * time.Second)
		}

		if err := models.FlushObjCache(&locl); err == nil {
			beego.Info("success set user cache")
			for true {
				time.Sleep(43200 * time.Second)
				models.FlushObjCache(&locl)
			}
		}else {
			beego.Error(err.Error())
		}
	}()

	beego.Debug("init")
}

func Login(username, password string) *User {
	sql := "select id, username, password from user where username = '" + username + "' and password = '" + password + "'"

	user, err := models.QueryFor(sql, &User{})

	if err == nil && len(user) > 0{
		u := user[0].Obj().(*User)
		return u
	}

	return nil
}
