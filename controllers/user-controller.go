package controllers

import (
	"fmt"
	"hello/models"

	"github.com/beego/beego/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {

	UserNameSpace = web.NewNamespace(
		"/user",

		web.NSPost(
			"/add", func(ctx *context.Context) {

				user := models.User{}

				if err := ctx.BindJSON(&user); err != nil {
					logs.Error(err)
					ctx.WriteString(err.Error())
				}

				ctx.JSONResp(user)
			},
		),

		web.NSGet(
			"/getOne", func(ctx *context.Context) {
				var id int
				if err := ctx.Input.Bind(&id, "id"); err != nil {
					logs.Error(err)
					ctx.WriteString(err.Error())
				}
				logs.Debug("---", id)
				ctx.WriteString(fmt.Sprint(id))
			},
		))

}

func logTest() {
	//an official log.Logger
	// l := logs.GetLogger()
	// l.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	// logs.GetLogger("ORM").Println("this is a message of orm")

	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")
}
