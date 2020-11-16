package sysInit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

var (
	bm cache.Cache
	isCacheInit bool
)

func init(){
	tbm, err := cache.NewCache("memory", `{"interval":60}`)
	bm = tbm
	if err == nil {
		beego.Debug("register cache success!")
		isCacheInit = true
	}else {
		beego.Error(err)
	}
}

func GetCache() cache.Cache{
	if bm != nil {
		return bm
	}
	beego.Error("bm is null")
	return nil
}

func IsCacheInit() bool{
	return isCacheInit
}
