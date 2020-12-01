package {{.packageName}}

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
	locl {{.upModelName}}
)

type {{.upModelName}}} struct {
    mu sync.Mutex
    Id  int64 `orm:"column(id);unique;pk"`{{range $i, $v := .paramMap}}
    {{$i}}  {{$v.type}}  `orm:"{{range $g := $v.prop}}{{$g}};{{end}}"`{{ end }}
}