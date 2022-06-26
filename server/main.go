package main

import (
	"fmt"
	"net"
)

// 处理与客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()
	// 读取客户端发送的信息
	for {
		infos := make([]byte, 1024*4)
		fmt.Println("读取客户端发送的数据...")
		n, err := conn.Read(infos[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read fail, err =", err)
			return
		}
		fmt.Println("读到的buf=", infos[:4])
	}
}

func main() {

	// 提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Listen fail, err =", err)
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
