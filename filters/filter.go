package filters

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	web.InsertFilter("/*", web.BeforeExec, filterFunc)
}

func filterFunc(ctx *context.Context) {

	// fmt.Println("lll", data)
	// type requestObjT struct{}
	// var requestObj = requestObjT{}
	// ctx.BindJSON(&requestObj)
	// logs.Debug("this is filterFunc", requestObj)
}
