package models

import(
    "core/models"
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
    orgId  int  `orm:"column(orgid);null;"`
    roleKey  string  `orm:"column(rolekey);unique;"`
    roleLevel  int  `orm:"column(rolelevel);"`
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

func (m *Role)ToModel() models.Model {
	return m.ToModel()
}


func (m *Role)QueryKey() string{
	return "id, name, orgid, rolekey, rolelevel"
}

func (m *Role)QueryResult(maps []orm.Params) ([]models.Model, error)  {
	if len(maps) == 0 {
		return nil, nil
	}

	mod := make([]Role, len(maps))
	models := make([]models.Model, len(maps))

	for i := 0; i < len(maps); i++ {
		id, err := strconv.ParseInt(fmt.Sprint(maps[i]["id"]), 10, 64)
		mod[i] = Role{
			Id: id,
			Name: fmt.Sprint(maps[i]["name"]),
			orgId: utils.ChangeValInt(fmt.Sprint(maps[i]["orgid"])),
			roleKey: fmt.Sprint(maps[i]["rolekey"]),
			roleLevel: utils.ChangeValInt(fmt.Sprint(maps[i]["rolelevel"])),
			}

		models[i] = mod[i].ToModel()
		if err != nil {
			return nil, err
		}
	}

	return models, nil
}

//实现String函数
func (m Role) String() string{
	return fmt.Sprint(
		`{"id": `, m.Id,
		    `, "name: "`, m.Name,
            `, "orgid: "`, m.orgId,
            `, "rolekey: "`, m.roleKey,
            `, "rolelevel: "`, m.roleLevel,
            `"}`)
}

func NewRole(m map[string]string) *Role {
	mod := Role{}

	id, err := strconv.ParseInt(m["id"], 10, 64)

	if err == nil {
		mod.Id = id
	}

	mod.Name = m["name"]
	mod.orgId = utils.ChangeValInt(m["orgid"]) 
	mod.roleKey = m["rolekey"]
	mod.roleLevel = utils.ChangeValInt(m["rolelevel"]) 
	
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

		if err := models.FlushObjCache(&loclRole); err == nil {
			beego.Info("success set role cache")
			for true {
				time.Sleep(43200 * time.Second)
				models.FlushObjCache(&loclRole)
			}
		}else {
			beego.Error(err.Error())
		}
	}()

	beego.Debug("init")
}