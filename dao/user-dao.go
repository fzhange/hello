package dao

import "hello/models"

func AddUser() {
	user := new(models.User)
	user.Age = 200

	OrmInstance.Insert(user)
}
