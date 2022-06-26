package util

import "fmt"

// 完成登录校验函数
func Login(userId int, userPwd string) error {
	// 定协议
	fmt.Printf("userId = %d, userPwd = %s\n", userId, userPwd)
	return nil
}
