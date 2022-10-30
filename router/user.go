package router

import (
	"hello/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	ns := web.NewNamespace(
		"/token",
		web.NSPost("/login", controllers.Login),
		web.NSPost("/exit", controllers.Exit),
		web.NSPost("/register", controllers.Register),
	)
	web.AddNamespace(ns)
}
