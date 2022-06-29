package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/client/utils"
	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/common/message"
)

// 显示登录成功后的界面
func ShowMenu() {
	fmt.Println("--------恭喜登录成功--------")
	fmt.Println("\t1.在线列表")
	fmt.Println("\t2.广播消息")
	fmt.Println("\t3.私密消息")
	fmt.Println("\t4.退出系统")
	fmt.Println("\t请选择(1~4):")

	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanln(&key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("广播发送：")
		fmt.Scanln(&content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("私密消息。。。")
	case 4:
		offline()
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
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		case message.UserStatusToOfflineType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			outputOfflineUser(&notifyUserStatusMes)
		default:
			fmt.Println("服务器返回了未知的消息类型")
		}
	}
}
