package clientProcess

import (
	"code/Project_practice/CommSystem_海量通讯系统/client/model"
	"code/Project_practice/CommSystem_海量通讯系统/common/message"
	"fmt"
	"time"
)

var clientUserOnline []string
var CurUser model.CurUser

func ShowOnlineUser() {
	fmt.Printf("当前在线用户(%v)\n", time.Now().Format("2006-Jan-02 03:04:05.999 pm"))
	for i := 0; i < len(clientUserOnline); i++ {

		fmt.Println("用户名：", clientUserOnline[i])
	}
}

func UpdateUserStatus(status *message.UserStatus) {
	clientUserOnline = append(clientUserOnline, status.AcoNum)
}