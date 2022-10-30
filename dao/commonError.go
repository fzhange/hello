package dao

import "errors"

var (
	ErrUserRepeat          = errors.New("user name has already exist")
	ErrUserNameIsEmpty     = errors.New("user name can't be empty")
	ErrUserPasswordIsEmpty = errors.New("user password can't be empty")
	ErrUserNotExist        = errors.New("<QuerySeter> no row found")
	ErrUserName            = errors.New("user name is invalid")
	ErrUserPassword        = errors.New("user password is invalid")
)
