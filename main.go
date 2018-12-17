package main

import (
	"github.com/astaxie/beego"
	_ "github.com/pengpan/beego_test/routers"
)

/*
You should execute the following script before running.

$ go get -u github.com/astaxie/beego
$ go get -u github.com/beego/bee
$ go get -u github.com/smartystreets/goconvey

@see https://beego.me/quickstart
*/
func main() {
	beego.Run()
}
