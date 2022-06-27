package process

import (
	"GoPlus/communication-system-zhuzi/server/utils"
	"fmt"
	"net"
	"os"
)

// 显示登录成功后的界面
func ShowMenu() {
	fmt.Println("--------恭喜登录成功--------")
	fmt.Println("\t1.在线列表")
	fmt.Println("\t2.发送消息")
	fmt.Println("\t3.信息列表")
	fmt.Println("\t4.退出系统")
	fmt.Println("\t请选择(1~4):")

	var key int
	fmt.Scanln(&key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("你退出了系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确")
	}
}

func ServerProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() fail, err =", err)
			return
		}
		fmt.Println("mes =", mes)
	}
}
