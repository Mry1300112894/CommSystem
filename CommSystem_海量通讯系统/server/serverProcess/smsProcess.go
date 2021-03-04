package serverProcess

import (
	"code/Project_practice/CommSystem_海量通讯系统/common/message"
	"code/Project_practice/CommSystem_海量通讯系统/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

func ForwardMes(mes *message.Message) (err error) {
	smsResMes := message.SmsMes{}
	// 反序列化 mes.Data
	err = json.Unmarshal([]byte(mes.Date), &smsResMes)
	if err != nil {
		fmt.Println("serverProcess json.Unmarshal([]byte(mes.Date), &smsResMes) err: ", err)
		return
	}

	// 将 mes 序列化
	data, err := json.Marshal(&mes)
	if err != nil {
		fmt.Println("serverProcess json.Marshal(mes) err: ", err)
		return
	}

	// 遍历在线用户列表
	for num, conn := range serverUserOnline {
		// 过滤自己，不会给自己发送消息
		if num == smsResMes.AcoNum {
			continue
		}

		// 否则, 将消息发送出去
		SendGroupMes(data, conn)
	}
	return
}

func SendGroupMes(data []byte, conn net.Conn){
	// 发送消息
	err := utils.WritePkg(conn, data)
	if err != nil {
		fmt.Println("serverProcess utils.WritePkg(conn, data) err: ", err)
		return
	}
}
