package services

import (
	"core/models"
	ssoMod "sso/models"
)

func Login(account, password string) *ssoMod.User {
	sql := "select id, username, password from user where account = '" + account + "' and password = '" + password + "'"

	user, err := models.QueryFor(sql, &ssoMod.User{})

	if err == nil && len(user) > 0{
		u := user[0].Obj().(*ssoMod.User)
		return u
	}

	return nil
}
