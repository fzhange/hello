package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

type CommonResponseStruct struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func CommonResponse(
	callback func(ctx *context.Context) (data interface{}, err error),
) func(ctx *context.Context) {

	return func(ctx *context.Context) {
		data, err := callback(ctx)
		logs.Info("CommonResponse err", err)
		logs.Info("CommonResponse data", data)
		if err != nil {
			ctx.JSONResp(
				CommonResponseStruct{
					Error: err.Error(),
					Data:  nil,
				},
			)
		} else {
			ctx.JSONResp(
				CommonResponseStruct{
					Error: "",
					Data:  data,
				},
			)
		}
	}
}
