package main

import (
	"github.com/beego/beego/logs"
	"github.com/beego/beego/v2/server/web"

	_ "hello/filters"

	// Delete this when you don't want to do test'
	// _ "hello/test"

	_ "hello/router"
	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	logs.EnableFuncCallDepth(true)
	web.BConfig.Listen.EnableAdmin = true
	// logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	// logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)

	logs.Info("web.BConfig.RunMode", web.BConfig.RunMode)
	if web.BConfig.RunMode == "dev" {
		// do some setting for dev mode
		// web.BConfig.WebConfig.DirectoryIndex = true
		// web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}

func main() {

	//  ---- beego start  ----
	web.Run()
	//  ---- beego end  ----

	// ---- demo start  ----
	// we should disable beego server, when we use demo program. beego server and demo server are conflict.
	// ---- demo end  ----
}
