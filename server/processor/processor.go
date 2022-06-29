package processor

import (
	"fmt"
	"io"
	"net"

	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/common/message"
	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/server/processes"
	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/server/utils"
)

type Processor struct {
	Conn net.Conn
}

// 根据客户端发送消息的种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	fmt.Println("mes =", mes)
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录逻辑
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册逻辑
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &processes.SmsProcess{}
		smsProcess.SendGrouppMes(mes)
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
