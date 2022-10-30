package main

import (
	"fmt"
	"time"

	"github.com/beego/beego/logs"
	"github.com/beego/beego/v2/server/web"

	"hello/controllers"
	_ "hello/filters"

	// Delete this when you don't want to do test'
	_ "hello/test"

	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	logs.EnableFuncCallDepth(true)
	// logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	// logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
}

func parseDemo() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006:01:02 15:04:05", "2022:10:05 11:25:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}

func main() {
	parseDemo()

	//  ---- beego start  ----
	web.BConfig.Listen.EnableAdmin = true
	// web.BConfig.Listen.AdminAddr = "localhost"
	// web.BConfig.Listen.AdminPort = 8088
	web.AddNamespace(controllers.UserNameSpace)
	web.Run()
	//  ---- beego end  ----

	// ---- demo start  ----
	// we should disable beego server, when we use demo program. beego server and demo server are conflict.
	// ---- demo end  ----
}
