module occ

go 1.15

require github.com/astaxie/beego v1.12.3

require github.com/smartystreets/goconvey v1.6.4

require (
	core v0.0.0
	github.com/OwnLocal/goes v1.0.0 // indirect
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726 // indirect
	github.com/siddontang/ledisdb v0.0.0-20181029004158-becf5f38d373 // indirect
	golang.org/x/tools v0.0.0-20200117065230-39095c1d176c // indirect
	sso v0.0.0
)

replace core => ../core

replace sso => ../sso
