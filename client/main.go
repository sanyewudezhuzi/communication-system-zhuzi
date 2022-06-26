package main

import (
	"GoPlus/communication-system-zhuzi/client/util"
	"fmt"
	"os"
)

var (
	userId  int
	userPwd string
)

func main() {

	// 接收用户的选择
	var key int
	// 判断是否还继续显示菜单
	var loop = true
	for loop {
		fmt.Println("--------多人聊天系统--------")
		fmt.Println("\t1.登录聊天室")
		fmt.Println("\t2.用户注册")
		fmt.Println("\t3.退出系统")
		fmt.Println("\t请选择(1~3):")
		fmt.Scanf("%d \n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("用户注册")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	// 根据用户的输出显示新的提示信息
	if key == 1 {
		// 说明用户要登录
		fmt.Println("请输入用户id：")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户pwd：")
		fmt.Scanln(&userPwd)
		// 先把登录的函数写到另外一个文件
		err := util.Login(userId, userPwd)
		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("进行注册用户。。")
	}

}
