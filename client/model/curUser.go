package model

import (
	"GoPlus/communication-system-zhuzi/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
