package dao

import (
	"hello/models"
)

func init() {

}

//~ insert operation
func InsertUsers(users []models.User) (int64, error) {
	return OrmInstance.InsertMulti(len(users), users)
}

func InserOneUser(user *models.User) (int64, error) {
	return OrmInstance.Insert(user)
}

//~ delete operation
func DeleteUser(user models.User, cols ...string) (int64, error) {
	return OrmInstance.Delete(user, cols...)
}

//~ get operation
func GetUserByField(user *models.User, cols ...string) error {
	return OrmInstance.Read(user, cols...)
}

//! report a (<QuerySeter> no row found) error when can't find a row in database
func GetUserByName(user *models.User, name string) error {
	return OrmInstance.Raw("SELECT * FROM user WHERE name = ?", name).QueryRow(user)
}

//~ update operation
func UpdateUser(user *models.User, cols ...string) (int64, error) {
	return OrmInstance.Update(user, cols...)
}
