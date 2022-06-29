package process

import (
	"GoPlus/communication-system-zhuzi/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &smsMes) fail, err =", err)
		return
	}

	info := fmt.Sprintf("用户id：\t%d 对大家说：\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}
