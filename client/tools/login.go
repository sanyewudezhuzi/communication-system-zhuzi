package tools

import (
	"GoPlus/communication-system-zhuzi/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// 完成登录校验函数
func Login(userId int, userPwd string) (err error) {

	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial(\"tcp\", \"127.0.0.1:8889\") fail, err =", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) fail, err =", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) fail, err =", err)
		return
	}

	// 根据规则 先发送长度 再发送内容
	pkgLen := uint32(len(data))
	var infos [4]byte
	binary.BigEndian.PutUint32(infos[:4], pkgLen)

	n, err := conn.Write(infos[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(infos[:4]) fail, err =", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail, err =", err)
		return
	}

	time.Sleep(10 * time.Second)
	fmt.Println("over..")

	return nil

}
