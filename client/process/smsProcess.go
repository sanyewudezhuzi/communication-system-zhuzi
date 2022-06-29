package process

import (
	"GoPlus/communication-system-zhuzi/common/message"
	"GoPlus/communication-system-zhuzi/server/utils"
	"encoding/json"
	"fmt"
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
