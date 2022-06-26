package main

import (
	"GoPlus/communication-system-zhuzi/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

// 读取数据
func readPkg(conn net.Conn) (mes message.Message, err error) {
	infos := make([]byte, 1024*4)
	fmt.Println("读取客户端发送的数据...")
	_, err = conn.Read(infos[:4])
	if err != nil {
		if err != io.EOF {
			fmt.Println("conn.Read(infos[:4]) fail, err =", err)
		}
		return
	}

	// 根据infos[:4]转成一共uint32类型
	pkgLen := binary.BigEndian.Uint32(infos[:4])
	n, err := conn.Read(infos[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(infos[:pkgLen]) fail, err =", err)
		return
	}

	err = json.Unmarshal(infos[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err =", err)
		return
	}
	return mes, nil

}

// 发送请求
func writePkg(conn net.Conn, data []byte) (err error) {

	pkgLen := uint32(len(data))
	var infos [4]byte
	binary.BigEndian.PutUint32(infos[:4], pkgLen)

	// 发送data长度
	n, err := conn.Write(infos[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(infos[:4]) fail, err =", err)
		return
	}

	// 发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail, err =", err)
		return
	}

	return
}

// 处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
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

	err = writePkg(conn, data)
	return

}

// 根据客户端发送消息的种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录逻辑
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
		// 处理注册逻辑
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

// 处理与客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()
	// 读取客户端发送的信息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端已退出")
			} else {
				fmt.Println("readPkg(conn) fail, err =", err)
			}
			return
		}
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

func main() {

	// 提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Listen(\"tcp\", \"127.0.0.1:8889\") fail, err =", err)
		return
	}
	defer listen.Close()

	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept fail, err =", err)
		}
		go process(conn)
	}
}
