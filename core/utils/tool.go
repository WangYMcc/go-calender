package utils

import "fmt"
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



