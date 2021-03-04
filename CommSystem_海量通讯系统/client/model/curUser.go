package model

import (
	"code/Project_practice/CommSystem_海量通讯系统/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
