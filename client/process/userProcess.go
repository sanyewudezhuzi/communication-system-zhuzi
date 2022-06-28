package process

import (
	"GoPlus/communication-system-zhuzi/common/message"
	"GoPlus/communication-system-zhuzi/server/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct{}

// 完成注册校验函数
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial(\"tcp\", \"127.0.0.1:8889\") fail, err =", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(registerMes) fail, err =", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) fail, err =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) fail, err =", err)
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("tf.ReadPkg() fail, err =", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功 请重新登录")
	} else {
		fmt.Println(registerResMes.Error)
	}
	os.Exit(0)
	return

}

// 完成登录校验函数
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

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

	fmt.Printf("客户端发送消息长度=%d 内容=%s", len(data), string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail, err =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("ReadPke() fail, err =", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {

		fmt.Println("当前在线用户列表：")
		for _, v := range loginResMes.UserIds {
			if v == userId {
				continue
			}
			fmt.Println("用户Id：\t", v)
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println("\n\n")

		// 起一个协程保持和服务器端的通讯，如果服务器有数据推送，则接收并显示再客户端的终端
		go ServerProcessMes(conn)

		for {
			ShowMenu()
		}

	} else {
		fmt.Println(loginResMes.Error)
	}
	return

}
