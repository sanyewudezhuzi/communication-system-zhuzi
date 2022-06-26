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
		fmt.Println("mes =", mes)
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
