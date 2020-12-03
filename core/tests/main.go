package main

import "github.com/astaxie/beego"

func main() {
	for i := 1; i < 30; i++ {
		objs := i
		size := 10
		page := 2
		beego.Info("total", i)
		beego.Info("Allpage", map[bool]int{true:1, false:0}[objs % size > 0]+ objs / size, ", nowPage", page)

		start := (page - 1) * size
		end := map[bool]int{true:page * size, false:objs % size}[objs >= page * size]

		beego.Info("start: ", start, ", end: ", end)
		beego.Debug("----------------------------------------------------------")
	}
}
