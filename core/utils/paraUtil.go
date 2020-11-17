package utils

import "fmt"

func CheckPkId(id int64) bool{
	if len(fmt.Sprint(id)) < 18{
		return false
	}

	return true
}
