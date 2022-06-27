package main

import (
	"GoPlus/communication-system-zhuzi/server/processor"
	"fmt"
	"net"
)

// 处理与客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()

	// 调用总控
	pro := &processor.Processor{
		Conn: conn,
	}
	err := pro.Process2()
	if err != nil {
		fmt.Println("processor.process2() fail, err =", err)
		return
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
