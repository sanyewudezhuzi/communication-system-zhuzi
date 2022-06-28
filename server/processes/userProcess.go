package processes

import (
	"GoPlus/communication-system-zhuzi/common/message"
	"GoPlus/communication-system-zhuzi/server/model"
	"GoPlus/communication-system-zhuzi/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn // 连接
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &registerMes) fail, err =", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(registerResMes) fail, err =", err)
		return
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail, err =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	// 先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginMes) fail, err =", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println("user =", user)
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) fail, err =", err)
		return
	}
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail, err =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}
