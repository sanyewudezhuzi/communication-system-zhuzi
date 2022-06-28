package main

import (
	"GoPlus/communication-system-zhuzi/client/process"
	"fmt"
	"os"
)

var (
	userId   int
	userPwd  string
	userName string
)

func main() {

	// 接收用户的选择
	var key int
	// 判断是否还继续显示菜单
	for {
		fmt.Println("--------多人聊天系统--------")
		fmt.Println("\t1.登录聊天室")
		fmt.Println("\t2.用户注册")
		fmt.Println("\t3.退出系统")
		fmt.Println("\t请选择(1~3):")
		fmt.Scanf("%d \n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id：")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户pwd：")
			fmt.Scanln(&userPwd)

			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("用户注册")
			fmt.Println("请输入注册用户id：")
			fmt.Scanln(&userId)
			fmt.Println("请输入注册用户pwd：")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入注册用户的name：")
			fmt.Scanln(&userName)

			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

}
