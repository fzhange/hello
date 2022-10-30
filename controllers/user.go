package controllers

import (
	"hello/dao"
	"hello/models"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {

	UserNameSpace = web.NewNamespace(
		"/token",

		web.NSPost(
			"/login", CommonResponse(func(ctx *context.Context) (data interface{}, err error) {
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
			}),
		),

		web.NSPost(
			"/exit", func(ctx *context.Context) {

				user := models.User{}

				ctx.JSONResp(user)
			},
		),

		web.NSPost(
			"/register", CommonResponse(func(ctx *context.Context) (interface{}, error) {

				requestUser := models.User{}

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
			}),
		),
	)
}
