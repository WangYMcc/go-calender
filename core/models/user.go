package models

import (
	"core/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"sync"
)

var (
	locl User
)

type User struct {
	mu sync.Mutex
	Id       int64 `orm:"column(id);unique;pk"`
	Username string `orm:"column(username)"`
	Password string `orm:"column(password);null"`
}

func (u *User) TableName() string {
	return "user"
}

func (u User) String() string{
	return fmt.Sprint("User=[id: ", strconv.FormatInt(u.Id, 10), ", username: ", u.Username, ", password: ", u.Password, "}")
}

func init(){
	orm.RegisterModel(new(User))
	locl = User{Id: 0}
	beego.Debug("init")
}

func SelectById(user User) (*User){
	o := orm.NewOrm()
	err := o.Read(&user)

	if err != nil {
		beego.Error(err)
		return nil
	}

	return &user
}

func SelectAll() ([]User){
	/*cache := sysInit.GetCache()
	if cache != nil && cache.Get("user") != nil {
		var user []User
		err := json.Unmarshal(cache.Get("user").([]uint8), &user)

		if err == nil {
			return user
		}

	}*/

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
	}

	return &user
}

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
		o.Commit()
	}

	return user
}

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

	return nil
}


