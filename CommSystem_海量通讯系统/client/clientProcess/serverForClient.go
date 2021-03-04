package clientProcess

import (
	"code/Project_practice/CommSystem_海量通讯系统/common/message"
	"code/Project_practice/CommSystem_海量通讯系统/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

// 和服务器保持通讯
func CommWithServer(conn net.Conn) {
	for {
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			fmt.Println("CommWithServer utils.ReadPkg(conn) err: ", err)
			return
		}
		switch mes.Type {
			// 如果发送的信息是这个类型，说明有新的用户上线了
			case message.UserStatusType:
				// 创建一个 UserStatus实例
				var userStatus message.UserStatus
				// 反序列化服务器发送来的消息
				err = json.Unmarshal([]byte(mes.Date), &userStatus)
				UpdateUserStatus(&userStatus)

			case message.SmsMesType:
				ReceiveGroupMes(&mes)
			default:
				fmt.Println("服务器返回了未知的消息！")
		}
	}
}



