package filters

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	web.InsertFilter("/*", web.BeforeExec, filterFunc)
}

func filterFunc(ctx *context.Context) {
	logs.Debug("this is filterFunc")
}
