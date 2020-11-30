package jwt

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateAccessToken(uid string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  uid,
		"exp":  time.Now().Add(time.Minute * 30).Unix(),
	})

	token, err := at.SignedString([]byte(beego.AppConfig.String("jwt.secret")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func CreateFlushToken(uid string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  uid,
		"exp":  time.Now().Add(time.Minute * 35).Unix(),
	})

	token, err := at.SignedString([]byte(beego.AppConfig.String("jwt.secret")))

	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(token string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwt.secret")), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["uid"].(string), nil
}
