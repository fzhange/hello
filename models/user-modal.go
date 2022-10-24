package models

// type User struct {
// 	ID  int `orm:"column(id)" json:"id"`
// 	Age int `orm:"column(age)" json:"age"`
// }

type User struct {
	ID  int `json:"id"`
	Age int `json:"age"`
}
