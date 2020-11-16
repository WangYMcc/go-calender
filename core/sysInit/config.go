package sysInit

import (
	"bufio"
	"github.com/astaxie/beego"
	"os"
	"strings"
)
var (
	config map[string]string
)
func init(){
	conf := make(map[string]string)

	f, err := os.Open("./conf/app.conf")   //因为bufio需要的是一个*os.File类型，所以我们换个方式读取，稍后再介绍一下
	if err != nil {
		beego.Error(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			beego.Error(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		keyval := s.Text()

		if strings.Contains(keyval, "."){
			arr := strings.Split(keyval, "=")
			conf[strings.TrimSpace(strings.Split(arr[0], ".")[1])] = strings.TrimSpace(arr[1])
		}else if keyval != "" {
			arr := strings.Split(keyval, "=")
			conf[strings.TrimSpace(arr[0])] = strings.TrimSpace(arr[1])
		}
	}

	err = s.Err()
	if err != nil {
		beego.Error(err)
	}

	config = conf

	beego.Debug("init")
}

func GetConf()  map[string]string {
	return config
}
