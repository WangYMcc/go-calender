package models

import (
	redis2 "core/sysInit/redis"
	"core/utils/snow"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
	"sync"
)

type Model interface {
	Obj() interface{}
	TableName() string
	GetLock() *sync.Mutex
	SetId(id int64)
	GetId() int64
	ToModel() Model
	QueryKey() string
	QueryResult(maps []orm.Params)  ([]Model, error)
}

func UpdateObjCache(obj Model) error{
	conn := redis2.GetRedisPool().Get()
	//注意close()
	defer conn.Close()

	o, err := json.Marshal(obj.Obj())
	conn.Do("SET", obj.TableName() + ":" +  fmt.Sprint(obj.GetId()), o)

	return err
}

func Insert(obj Model) interface{} {
	workderc, _ := snow.NewWorker(int64(1))
	id := workderc.GetId()
	obj.SetId(id)

	obj.GetLock().Lock()
	defer obj.GetLock().Unlock()

	o := orm.NewOrm()
	o.Begin()

	_, err := o.Insert(obj.Obj())

	if err != nil {
		beego.Error(err)
		o.Rollback()
		return nil
	}else {
		o.Commit()
		obj.SetId(id)

		err := UpdateObjCache(obj)
		if err != nil {
			beego.Error(err)
		}
	}

	return obj.Obj()
}

func InsertMore(objs []Model) ([]Model, error){
	if len(objs) == 0 {
		return nil, fmt.Errorf("input is null")
	}

		workderc, _ := snow.NewWorker(int64(1))
	var err error

	o := orm.NewOrm()
	o.Begin()

	ids := make([]int64, len(objs))
	for i := 0; i < len(objs); i++ {
		ids[i] = workderc.GetId()

		objs[i].SetId(ids[i])
		if err == nil {
			_, err = o.Insert(objs[i].Obj())
		}else {
			break
		}
	}

	if err != nil {
		o.Rollback()
		return nil, err
	}else{
		for i := 0; i < len(objs); i++ {
			objs[i].SetId(ids[i])
			UpdateObjCache(objs[i])
		}
		o.Commit()
	}

	return objs, nil
}

func FuzzyQuery(key string, value string, obj Model) ([]Model, error) {
	var maps []orm.Params

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(obj.QueryKey()).From(obj.TableName()).Where(key + " like '%" + value + "%'")
	sql := qb.String()

	o := orm.NewOrm()
	_, err := o.Raw(sql).Values(&maps)

	if err != nil {
		return nil, err
	}

	return obj.QueryResult(maps)
}

func AccurateQuery(key string, value string, obj Model) ([]Model, error) {
	var maps []orm.Params

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(obj.QueryKey()).From(obj.TableName()).Where(key + " = '"+ value + "'")
	sql := qb.String()

	o := orm.NewOrm()
	_, err := o.Raw(sql).Values(&maps)

	if err != nil {
		return nil, err
	}

	return obj.QueryResult(maps)
}

func QueryFor(sql string, obj Model) ([]Model, error) {
	var maps []orm.Params

	o := orm.NewOrm()
	_, err := o.Raw(sql).Values(&maps)

	if err != nil {
		return nil, err
	}

	return obj.QueryResult(maps)
}

func QueryForMap(sql string) ([]orm.Params, error) {
	var maps []orm.Params

	o := orm.NewOrm()
	_, err := o.Raw(sql).Values(&maps)

	if err != nil {
		return nil, err
	}

	return maps, nil
}

//刷新redis 某个表缓存
func FlushObjCache(obj Model) error{
	conn := redis2.GetRedisPool().Get()
	//注意close()
	defer conn.Close()

	arrs, err_ := SelectAll(obj)

	if err_ != nil {
		return err_
	}

	var u []byte
	var err error

	for i := 0; i < len(arrs); i++ {
		u, err = json.Marshal(arrs[i].Obj())
		conn.Do("SET", obj.TableName() + ":" +  fmt.Sprint(arrs[i].GetId()), u)
	}

	return err
}

//从数据库中获取所有user信息
func SelectAll(obj Model) ([]Model, error){
	var maps []orm.Params

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(obj.QueryKey()).From(obj.TableName())
	sql := qb.String()

	o := orm.NewOrm()
	_, err := o.Raw(sql).Values(&maps)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return obj.QueryResult(maps)
}

//从数据库中通过id获取表信息
func SelectById(obj Model) (Model){
	conn := redis2.GetRedisPool().Get()
	defer conn.Close()

	result, _ := conn.Do("GET", obj.TableName() + ":" + fmt.Sprint(obj.GetId()))
	if result != nil {
		bytes, _ := redis.Bytes(result, nil)
		err := json.Unmarshal(bytes, obj.Obj())
		if err == nil {
			return obj
		}
	}

	o := orm.NewOrm()
	err := o.Read(obj.Obj())
	if err != nil {
		beego.Error(err)
		return nil
	}

	data, _ := json.Marshal(obj.Obj())

	conn.Do("SET", obj.TableName() + ":" + fmt.Sprint(obj.GetId()), data)
	return obj
}

//批量更新数据库表信息
func UpdateMore(objs []Model) ([]Model, error){
	var err error

	o := orm.NewOrm()
	o.Begin()
	for i := 0; i < len(objs); i++ {
		if err == nil {
			_, err = o.Update(objs[i].Obj())
		}else {
			break
		}
	}

	if err != nil {
		o.Rollback()
		return []Model{}, err
	}else{
		for i := 0; i < len(objs); i++ {
			UpdateObjCache(objs[i])
		}
		o.Commit()
	}

	return objs, nil
}

//更新数据库表单条信息
func Update(obj Model) (error){
	obj.GetLock().Lock()
	defer obj.GetLock().Unlock()

	o := orm.NewOrm()
	o.Begin()

	n, err := o.Update(obj.Obj())
	if err != nil{
		o.Rollback()
		beego.Error(err)
		return err
	}else if n == 0 {
		o.Rollback()
		RemoveObjCache(obj)
		return fmt.Errorf("not such item exist")
	}

	o.Commit()
	UpdateObjCache(obj)

	return nil
}

func Delete(obj Model) (error) {
	obj.GetLock().Lock()
	defer obj.GetLock().Unlock()

	o := orm.NewOrm()
	o.Begin()
	_, err := o.Delete(obj.Obj())

	if err == nil {
		RemoveObjCache(obj)
		o.Commit()
	}else {
		o.Rollback()
	}

	return err
}

func DeleteMore(objs []Model) (error) {
	o := orm.NewOrm()
	o.Begin()

	var err_o error
	for i := 0; i < len(objs); i++ {
		_, err := o.Delete(objs[i].Obj())
		if err != nil {
			err_o = err
			break
		}
	}

	if err_o == nil {
		o.Commit()
		for i := 0; i < len(objs); i++ {
			RemoveObjCache(objs[i])
		}
	}else {
		o.Rollback()
	}

	return err_o
}

//移除redis单个缓存
func RemoveObjCache(obj Model) error{
	conn := redis2.GetRedisPool().Get()
	//注意close()
	defer conn.Close()

	_, err := conn.Do("DEL", obj.TableName() + ":" + fmt.Sprint(obj.GetId()))

	return err
}

func SqlPage(sql string, pageNum int, pageSize int) string{
	sql += " limit " + fmt.Sprint(pageNum - 1) + ", " + fmt.Sprint(pageSize)

	return sql
}

func Count(obj Model) (int64, error){
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("count(id) as count").From(obj.TableName())
	sql := qb.String()

	o := orm.NewOrm()
	res, err := o.Raw(sql).Exec()

	if err == nil {
		num, err_ := res.RowsAffected()

		if err_ == nil {
			return num, nil
			err = err_
		}
	}

	return 0, err
}

func AllPage(obj Model, pageSize int64) (int64, error){
	num, err := Count(obj)

	if err == nil {
		n := map[bool]int64{true: 1, false: 0}[num % pageSize > 0]
		return (num / pageSize) + n, nil
	}

	return 0, nil
}