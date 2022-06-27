package processes

import (
	"GoPlus/communication-system-zhuzi/common/message"
	"GoPlus/communication-system-zhuzi/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn // 连接
}

// 处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	// 先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginMes) fail, err =", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	// id = 100 pwd = 123456
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在"
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) fail, err =", err)
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail, err =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}
