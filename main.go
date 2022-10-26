package main

import (
	"hello/demo"

	"github.com/beego/beego/logs"

	// don't forget this

	_ "hello/filters"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	logs.EnableFuncCallDepth(true)
	// logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	// logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
}

func main() {
	//  ---- beego start  ----
	// web.BConfig.Listen.EnableAdmin = true
	// web.BConfig.Listen.AdminAddr = "localhost"
	// web.BConfig.Listen.AdminPort = 8088
	// web.AddNamespace(controllers.UserNameSpace)
	// web.Run()
	//  ---- beego end  ----

	// ---- demo start  ----
	// we should disable beego server, when we use demo program. beego server and demo server are conflict.
	demo.StartServer()
	// ---- demo end  ----
	var a, b int
	if a, ok = b.(int);  {

	}
}
