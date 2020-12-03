package utils

import (
	"core/models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)
var (
	HTTP_OK = 200
	HTTP_NOT_FOUND = 404
	HTTP_USER_ERROR = 400
	HTTP_SERVER_ERROR = 500
)

type Result struct {
	Total int  //总条目
	CurSize int  //当前页条目
	Page int //当前页
	AllPage int  //页数
	Objs interface{} //数组
}

type MessageResponse struct {
	Code int
	Message string
	List interface{}
	Obj interface{}
}

func CheckPkId(id int64) bool{
	if len(fmt.Sprint(id)) < 18{
		return false
	}

	return true
}

func GenerateRequest(code int, message string) MessageResponse {
	return MessageResponse{Code: code, Message: message}
}

func GenerateRequestWithObj(code int, message string, result interface{}) MessageResponse {
	return MessageResponse{Code: code, Message: message, Obj: result}
}

func GenerateRequestWithList(code int, message string, result Result) MessageResponse {
	return MessageResponse{Code: code, Message: message, List: result}
}

func ChangeValInt64(val string) int64{
	v, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		beego.Error(err)
		return -1
	}

	return v
}
func ChangeValInt(val string) int{
	v, err := strconv.Atoi(val)

	if err != nil {
		beego.Error(err)
		return -1
	}

	return v
}

func StartPage(page int, size int, objs []models.Model) Result{
	result := Result{Page: page, Total: len(objs), AllPage: map[bool]int{true:1, false:0}[len(objs) % size > 0] + (len(objs) / size), CurSize: size}
	start := 0
	end := 0

	if result.Page > result.AllPage {
		result.Objs = nil
	}else if result.Page == result.AllPage {
		start = (page - 1) * size
		end = map[bool]int{true:page * size, false:len(objs) % size + start}[len(objs) >= page * size]
	}else {
		start = (page - 1) * size
		end = map[bool]int{true:page * size, false:len(objs) % size + start}[len(objs) >= page * size]
	}

	if start < end {
		curObjs := make([]interface{}, end - start)
		index := 0

		for i := start; i < end; i++ {
			curObjs[index] = objs[i]
			index++
		}

		result.Objs = curObjs
	}

	return result
}


