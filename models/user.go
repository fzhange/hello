package models

type User struct {
	Id       int    `json:"id"` //int, int32 - 设置 auto 或者名称为 Id 时
	Name     string `json:"name"`
	Password string `json:"password"`
}
