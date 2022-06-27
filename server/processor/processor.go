package processor

import (
	"GoPlus/communication-system-zhuzi/common/message"
	"GoPlus/communication-system-zhuzi/server/processes"
	"GoPlus/communication-system-zhuzi/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 根据客户端发送消息的种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录逻辑
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册逻辑
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) Process2() (err error) {
	// 读取客户端发送的信息
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端已退出")
			} else {
				fmt.Println("readPkg(conn) fail, err =", err)
			}
			return err
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
