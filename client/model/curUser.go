package model

import (
	"net"

	"github.com/NotAPigInTheTrefoilHouse/communication-system-zhuzi/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
