package {{.modelPackageName}}

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
	locl{{.upModelName}} {{.upModelName}}
)

type {{.upModelName}} struct {
    mu sync.Mutex
    Id  int64 `orm:"column(id);unique;pk"`{{range $i, $v := .paramMap}}
    {{$i}}  {{$v.type}}  `orm:"{{range $g := $v.prop}}{{$g}};{{end}}"`{{ end }}
}

func (m *{{.upModelName}})Obj() interface{} {
	return m
}

func (m *{{.upModelName}})SetId(id int64){
	m.Id = id
}

func (m *{{.upModelName}})GetLock() *sync.Mutex{
	return &m.mu
}

//设置表名
func (m *{{.upModelName}}) TableName() string {
	return "{{.tableName}}"
}

func (m *{{.upModelName}})GetId() int64{
	return m.Id
}

func (m *{{.upModelName}})ToModel() models.Model {
	return m.ToModel()
}


func (m *{{.upModelName}})QueryKey() string{
	return "id{{range $i, $v := .paramMap}}, {{$v.low}}{{end}}"
}

func (m *{{.upModelName}})QueryResult(maps []orm.Params) ([]models.Model, error)  {
	if len(maps) == 0 {
		return nil, nil
	}

	mod := make([]{{.upModelName}}, len(maps))
	models := make([]models.Model, len(maps))

	for i := 0; i &lt; len(maps); i++ {
		id, err := strconv.ParseInt(fmt.Sprint(maps[i]["id"]), 10, 64)
		mod[i] = {{.upModelName}}{
			Id: id,
			{{range $i, $v := .paramMap}}{{$i}}: {{if $v.int}}utils.ChangeValInt(fmt.Sprint(maps[i]["{{$v.low}}"])){{else}}fmt.Sprint(maps[i]["{{$v.low}}"]){{end}},
			{{end}}}

		models[i] = mod[i].ToModel()
		if err != nil {
			return nil, err
		}
	}

	return models, nil
}

//实现String函数
func (m {{.upModelName}}) String() string{
	return fmt.Sprint(
		`{"id": `, m.Id,
		    {{range $i, $v := .paramMap}}`, "{{$v.low}}: "`, m.{{$i}},
            {{end}}`"}`)
}

func New{{.upModelName}}(m map[string]string) *{{.upModelName}} {
	mod := {{.upModelName}}{}

	id, err := strconv.ParseInt(m["id"], 10, 64)

	if err == nil {
		mod.Id = id
	}

	{{range $i, $v := .paramMap}}mod.{{$i}} = {{if $v.int}}utils.ChangeValInt(m["{{$v.low}}"]) {{else}}m["{{$v.low}}"]{{end}}
	{{end}}
	return &mod
}

//初始化-建表-将数据存到缓存中
func init(){
	orm.RegisterModel(new({{.upModelName}}))
	locl{{.upModelName}} = {{.upModelName}}{Id: 0}

	go func(){
		for !sql.IsOrmInit() {
			time.Sleep(1 * time.Second)
		}

		sql.RunSyncDb()

		for !redis.IsRedisCacheInit() {
			time.Sleep(1 * time.Second)
		}

		if err := models.FlushObjCache(&locl{{.upModelName}}); err == nil {
			beego.Info("success set {{.tableName}} cache")
			for true {
				time.Sleep(43200 * time.Second)
				models.FlushObjCache(&locl{{.upModelName}})
			}
		}else {
			beego.Error(err.Error())
		}
	}()

	beego.Debug("init")
}