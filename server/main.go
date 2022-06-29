package main

import (
	"fmt"
	"net"
	"time"

	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/server/model"
	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/server/processor"
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

func initUserDao() {
	model.MyUserDao = model.NewUserDao(processor.Pool)
}

func main() {

	processor.InitPool("127.0.0.1:6379", 16, 0, 300*time.Second)
	initUserDao()
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
