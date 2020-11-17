package models

import (
	"core/sysInit"
	"core/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
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
	Username string `orm:"column(username)"`
	Password string `orm:"column(password);null"`
}

//设置表名
func (u *User) TableName() string {
	return "user"
}

//实现String函数
func (u User) String() string{
	return fmt.Sprint("User=[id: ", strconv.FormatInt(u.Id, 10), ", username: ", u.Username, ", password: ", u.Password, "}")
}

//初始化-建表-将数据存到缓存中
func init(){
	orm.RegisterModel(new(User))
	locl = User{Id: 0}

	go func(){
		for !sysInit.IsOrmInit() {
			time.Sleep(1 * time.Second)
		}

		sysInit.RunSyncDb()

		for !sysInit.IsRedisCacheInit() {
			time.Sleep(1 * time.Second)
		}

		FlushUserCache();

		beego.Info("success set user cache")
	}()

	beego.Debug("init")
}

//刷新redis所有user缓存
func FlushUserCache() error{
	conn := sysInit.GetRedisPool().Get()
	//注意close()
	defer conn.Close()
	user := SelectAll()
	var err error
	var u []byte

	for i := 0; i < len(user); i++ {
		u, err = json.Marshal(user[i])
		conn.Do("SET", "user:" +  fmt.Sprint(user[i].Id), u)
	}

	return err
}

//更新redis单个user缓存
func UpdateUserCache(user User) error{
	conn := sysInit.GetRedisPool().Get()
	//注意close()
	defer conn.Close()

	u, err := json.Marshal(user)
	conn.Do("SET", "user:" +  fmt.Sprint(user.Id), u)

	return err
}

//移除redis单个user缓存
func RemoveUserCache(id int64) error{
	conn := sysInit.GetRedisPool().Get()
	//注意close()
	defer conn.Close()

	_, err := conn.Do("DEL", "user:" + fmt.Sprint(id))

	return err
}

//从数据库中通过id获取user信息
func SelectById(id int) (*User){
	var user User
	conn := sysInit.GetRedisPool().Get()
	defer conn.Close()

	result, _ := conn.Do("GET", "user:" + string(id))
	if result != nil {
		bytes, _ := redis.Bytes(result, nil)
		err := json.Unmarshal(bytes, &user)
		if err == nil {
			return &user
		}
	}

	o := orm.NewOrm()
	err := o.Read(&user)

	if err != nil {
		beego.Error(err)
		return nil
	}

	data, _ := json.Marshal(user)

	conn.Do("SET", "user:" + string(id), data)
	return &user
}

//从数据库中获取所有user信息
func SelectAll() ([]User){
	var user []User

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id", "username", "password").From("user")
	sql := qb.String()

	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&user)

	if err != nil {
		beego.Error(err)
		return nil
	}

	return user
}

//从数据库中获取分页的user信息
func SelectByPage(displayNum int, page int) ([]User){
	var user []User

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id", "username", "password").From("user").Limit(displayNum).Offset((page - 1) * displayNum)
	sql := qb.String()

	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&user)

	if err != nil {
		beego.Error(err)
		return nil
	}

	return user
}

//添加一条user信息
func (user User)Insert() (*User){
	workderc, _ := utils.NewWorker(int64(1))
	user.Id = workderc.GetId()

	user.mu.Lock()
	defer user.mu.Unlock()

	o := orm.NewOrm()
	o.Begin()

	_, err := o.Insert(&user)

	if err != nil {
		beego.Error(err)
		o.Rollback()
		return nil
	}else {
		o.Commit()
		err := UpdateUserCache(user)

		if err != nil {
			beego.Error(err)
		}
	}

	return &user
}

//添加大量user信息
func InsertMore(user []User) ([]User){
	workderc, _ := utils.NewWorker(int64(1))
	var err error

	o := orm.NewOrm()
	o.Begin()

	for i := 0; i < len(user); i++ {
		user[i].Id = workderc.GetId()
		if err == nil {
			_, err = o.Insert(&user[i])
		}else {
			break
		}
	}

	if err != nil {
		o.Rollback()
	}else{
		for i := 0; i < len(user); i++ {
			UpdateUserCache(user[i])
		}
		o.Commit()
	}

	return user
}

//更新条user信息
func (user User) Update() (error){
	user.mu.Lock()
	defer user.mu.Unlock()

	o := orm.NewOrm()
	o.Begin()

	_, err := o.Update(&user)

	if err != nil {
		o.Rollback()
		beego.Error(err)
		return err
	}
	o.Commit()
	UpdateUserCache(user)

	return nil
}

func (user User) Delete() (error) {
	user.mu.Lock()
	defer user.mu.Unlock()

	o := orm.NewOrm()
	o.Begin()
	_, err := o.Delete(&user)

	if err == nil {
		RemoveUserCache(user.Id)
		o.Commit()
	}else {
		o.Rollback()
	}

	return err
}

func DeleteMore(users []User) (error) {
	locl.mu.Lock()
	defer locl.mu.Unlock()

	o := orm.NewOrm()
	o.Begin()

	var err_o error
	for i := 0; i < len(users); i++ {
		_, err := o.Delete(&users[i])
		if err != nil {
			err_o = err
			break
		}

	}

	if err_o == nil {
		o.Commit()
		for i := 0; i < len(users); i++ {
			RemoveUserCache(users[i].Id)
		}
	}else {
		o.Rollback()
	}

	return err_o
}


