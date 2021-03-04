package main

import (
	"code/Project_practice/CommSystem_海量通讯系统/client/clientProcess"
	"fmt"
	"net"
)

func main() {
	// 和服务器建立连接
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial() err: ", err)
		return
	}
	// 延时关闭连接
	defer conn.Close()

	clientProcess.View1(conn)
}