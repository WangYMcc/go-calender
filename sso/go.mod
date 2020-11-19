module sso

go 1.15

require github.com/astaxie/beego v1.12.3

require (
	core v0.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/smartystreets/goconvey v1.6.4
)

replace core => ../core
