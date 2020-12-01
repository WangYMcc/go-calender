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
	Name string `orm:"column(name)"`
	Account string `orm:"column(account);unique"`
	Password string `orm:"column(password)"`
	Phone string `orm:"column(phone);null"`
	Email string `orm:"column(email);null"`
	Sex string `orm:"column(sex);null"`
	UserKey string `orm:"column(userkey);null"`
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
	return "id, name, account, password, phone, email, sex, userkey"
}

func (u *User)QueryResult(maps []orm.Params) ([]models.Model, error)  {
	if len(maps) == 0 {
		return nil, nil
	}

	user := make([]User, len(maps))
	models := make([]models.Model, len(maps))

	for i := 0; i < len(maps); i++ {
		id, err := strconv.ParseInt(fmt.Sprint(maps[i]["id"]), 10, 64)
		user[i] = User{
			Id: id,
			Name: fmt.Sprint(maps[i]["name"]),
			Password: fmt.Sprint(maps[i]["password"]),
			Account: fmt.Sprint(maps[i]["account"]),
			Phone: fmt.Sprint(maps[i]["phone"]),
			Email: fmt.Sprint(maps[i]["email"]),
			Sex: fmt.Sprint(maps[i]["sex"]),
			UserKey: fmt.Sprint(maps[i]["userkey"])}

		models[i] = user[i].ToModel()
		if err != nil {
			return nil, err
		}
	}

	return models, nil
}

//实现String函数
func (u User) String() string{
	return fmt.Sprint(
		`{"id": `, u.Id,
			`, "name": "`, u.Name,
			`", "account": "`, u.Account,
			`", "password": "`, u.Password,
			`", "phone": "`, u.Phone,
			`", "email": "`, u.Email,
			`", "sex": "`, u.Sex,
			`", "userkey": "`, u.UserKey,
		`"}`)
}

func NewUser(m map[string]string) *User {
	user := User{}

	id, err := strconv.ParseInt(m["id"], 10, 64)

	if err == nil {
		user.Id = id
	}
	user.Name = m["name"]
	user.Account = m["account"]
	user.Password = m["password"]
	user.Phone = m["phone"]
	user.Email = m["email"]
	user.Sex = m["sex"]
	user.UserKey = m["userkey"]

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