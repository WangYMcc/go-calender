package models

import(
    coreMod "core/models"
    "core/utils"
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
	loclRole Role
)

type Role struct {
    mu sync.Mutex
    Id  int64 `orm:"column(id);unique;pk"`
    Name  string  `orm:"column(name);"`
    OrgKey  int  `orm:"column(orgkey);null;"`
    RoleKey  string  `orm:"column(rolekey);unique;"`
    RoleLevel  int  `orm:"column(rolelevel);"`
}

func (m *Role)Obj() interface{} {
	return m
}

func (m *Role)SetId(id int64){
	m.Id = id
}

func (m *Role)GetLock() *sync.Mutex{
	return &m.mu
}

//设置表名
func (m *Role) TableName() string {
	return "role"
}

func (m *Role)GetId() int64{
	return m.Id
}

func (m *Role)QueryKey() string{
	return "id, name, orgkey, rolekey, rolelevel"
}

func (m *Role)QueryResult(maps []orm.Params) ([]coreMod.Model, error)  {
	if len(maps) == 0 {
		return nil, nil
	}

	mod := make([]Role, len(maps))
	coreMod := make([]coreMod.Model, len(maps))

	for i := 0; i < len(maps); i++ {
		id, err := strconv.ParseInt(fmt.Sprint(maps[i]["id"]), 10, 64)
		mod[i] = Role{
			Id: id,
			Name: fmt.Sprint(maps[i]["name"]),
			OrgKey: utils.ChangeValInt(fmt.Sprint(maps[i]["orgkey"])),
			RoleKey: fmt.Sprint(maps[i]["rolekey"]),
			RoleLevel: utils.ChangeValInt(fmt.Sprint(maps[i]["rolelevel"])),
			}

		coreMod[i] = &mod[i]
		if err != nil {
			return nil, err
		}
	}

	return coreMod, nil
}

//实现String函数
func (m Role) String() string{
	return fmt.Sprint(
		`{"id": `, m.Id,
		    `, "name: "`, m.Name,
            `, "orgkey: "`, m.OrgKey,
            `, "rolekey: "`, m.RoleKey,
            `, "rolelevel: "`, m.RoleLevel,
            `"}`)
}

func NewRole(m map[string]string) *Role {
	mod := Role{}

	id, err := strconv.ParseInt(m["id"], 10, 64)

	if err == nil {
		mod.Id = id
	}

	mod.Name = m["name"]
	mod.OrgKey = utils.ChangeValInt(m["orgkey"]) 
	mod.RoleKey = m["rolekey"]
	mod.RoleLevel = utils.ChangeValInt(m["rolelevel"]) 
	
	return &mod
}

//初始化-建表-将数据存到缓存中
func init(){
	orm.RegisterModel(new(Role))
	loclRole = Role{Id: 0}

	go func(){
		for !sql.IsOrmInit() {
			time.Sleep(1 * time.Second)
		}

        sql.RunSyncDb()

		for !redis.IsRedisCacheInit() {
			time.Sleep(1 * time.Second)
		}

		if err := coreMod.FlushObjCache(&loclRole); err == nil {
			beego.Info("success set role cache")
			for true {
				time.Sleep(43200 * time.Second)
				coreMod.FlushObjCache(&loclRole)
			}
		}else {
			beego.Error(err.Error())
		}
	}()

	beego.Debug("init")
}