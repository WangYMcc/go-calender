package models

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetConf(val string)  map[string]string {
	conf := make(map[string]string)

	f, err := os.Open("./conf/app.conf")   //因为bufio需要的是一个*os.File类型，所以我们换个方式读取，稍后再介绍一下
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		keyval := s.Text()

		if strings.Contains(keyval, val) {
			arr := strings.Split(keyval, "=")
			conf[strings.TrimSpace(strings.Split(arr[0], ".")[1])] = strings.TrimSpace(arr[1])
		}

	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	return conf
}
