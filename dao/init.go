package dao

import (
	"hello/models"

	"github.com/beego/beego/v2/client/orm"
	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

var OrmInstance orm.Ormer

func init() {
	// need to register models in init
	orm.RegisterModel(new(models.User))

	// need to register default database
	orm.RegisterDataBase("default", "mysql", "root:Pa123457@tcp(127.0.0.1:3306)/ssmp?charset=utf8")

	OrmInstance = orm.NewOrm()
	// automatically build table
	orm.RunSyncdb("default", false, true)

}
