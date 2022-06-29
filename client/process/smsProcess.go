package process

import (
	"encoding/json"
	"fmt"

	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/client/utils"
	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/common/message"
)

type SmsProcess struct{}

func (this *SmsProcess) SendGroupMes(content string) (err error) {

	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) fail, err =", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) fail, err =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) fail, err =", err)
	}
	return

}
