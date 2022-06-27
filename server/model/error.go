package model

import "errors"

var (
	ERROR_USER_NOTEXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS    = errors.New("用户已被注册")
	ERROR_USER_PWD       = errors.New("密码错误")
)
