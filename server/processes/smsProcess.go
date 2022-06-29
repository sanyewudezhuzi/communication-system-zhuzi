package processes

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/common/message"
	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/server/utils"
)

type SmsProcess struct{}

func (this *SmsProcess) SendGrouppMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &smsMes) fail, err =", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) fail, err =", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) fail, err =", err)
	}
}
