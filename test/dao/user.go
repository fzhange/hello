package dao

import (
	"hello/dao"
	"hello/models"

	"github.com/beego/beego/logs"
)

func init() {
	getUserByName()
}

func getUserByName() {
	user := models.User{}
	err := dao.GetUserByName(&user, "tony")
	logs.Debug("getUserByName test", user, err)
}
