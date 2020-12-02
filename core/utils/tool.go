package utils

import (
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

type MessageResponse struct {
	Code int
	Message string
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

func ChangeValInt64(val string) int64{
	v, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		beego.Error(err)
		return 0
	}

	return v
}
func ChangeValInt(val string) int{
	v, err := strconv.Atoi(val)

	if err != nil {
		beego.Error(err)
		return 0
	}

	return v
}


