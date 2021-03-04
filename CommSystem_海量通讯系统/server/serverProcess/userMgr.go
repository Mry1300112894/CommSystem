package serverProcess

import (
	"fmt"
	"net"
)

// 搞一个 map 用来存放在线用户
var serverUserOnline = make(map[string]net.Conn, 1024)

// 添加用户到列表中
func AddUser(acoNum string, conn net.Conn) {
	serverUserOnline[acoNum] = conn
}

// 删除用户
func DelUser(acoNum string) {
	delete(serverUserOnline, acoNum)
}

// 获取所有用户
func GetAllUser() map[string]net.Conn {
	return serverUserOnline
}

// 获取单个用户
func GetAnyUser(acoNum string) (conn net.Conn, err error) {
	conn, ok := serverUserOnline[acoNum]
	if !ok {
		err = fmt.Errorf("用户：%v不存在！", acoNum)
		return
	}
	return
}
