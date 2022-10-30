package controllers

import (
	"hello/dao"
	"hello/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
)

func Login(ctx *context.Context) {
	CommonResponse(func(ctx *context.Context) (data interface{}, err error) {
		requestUser := models.User{}

		if err := ctx.BindJSON(&requestUser); err != nil {
			return nil, err
		}

		dbUser := models.User{}
		err = dao.GetUserByName(&dbUser, requestUser.Name)

		if err != nil && err.Error() == dao.ErrUserNotExist.Error() {
			return nil, dao.ErrUserNotExist
		}

		if dbUser.Password != requestUser.Password {
			return nil, dao.ErrUserPassword
		}

		return "login success", nil
	})(ctx)
}

func Exit(ctx *context.Context) {
	user := models.User{}

	ctx.JSONResp(user)
}

func Register(ctx *context.Context) {
	CommonResponse(func(ctx *context.Context) (interface{}, error) {

		requestUser := models.User{}

		logs.Info("Registering", requestUser)

		if err := ctx.BindJSON(&requestUser); err != nil {
			return nil, err
		}

		if requestUser.Name == "" {
			return nil, dao.ErrUserNameIsEmpty
		}

		if requestUser.Password == "" {
			return nil, dao.ErrUserPasswordIsEmpty
		}

		dbUser := models.User{}
		err := dao.GetUserByName(&dbUser, requestUser.Name)
		if err != nil && err.Error() == dao.ErrUserNotExist.Error() {
			pkId, err := dao.InserOneUser(&requestUser)
			if err == nil {
				return pkId, nil
			} else {
				return nil, err
			}
		} else {
			return nil, dao.ErrUserRepeat
		}
	})(ctx)
}
