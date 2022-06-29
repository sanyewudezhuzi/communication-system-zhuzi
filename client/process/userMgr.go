package process

import (
	"fmt"

	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/client/model"
	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/common/message"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

// 在客户端显示当前在线用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id := range onlineUsers {
		fmt.Println("用户id =：\t", id)
	}
}

// 处理返回信息
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}

// 离线
func outputOfflineUser(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	notifyUserStatusMes.Status = message.UserOffline
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
func offline() {

	var mes message.Message
	mes.Type = message.LoginMesType
}
