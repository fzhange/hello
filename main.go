package main

import (
	"fmt"
	"hello/controllers"

	"github.com/beego/beego/logs"
	"github.com/beego/beego/v2/server/web"

	// don't forget this

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	logs.EnableFuncCallDepth(true)
	// logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	// logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
}

func main() {
	web.BConfig.Listen.EnableAdmin = true
	// web.BConfig.Listen.AdminAddr = "localhost"
	// web.BConfig.Listen.AdminPort = 8088

	web.AddNamespace(controllers.UserNameSpace)

	i := 10
	fmt.Printf("%d %T \n", i, i)

	web.Run()

}
