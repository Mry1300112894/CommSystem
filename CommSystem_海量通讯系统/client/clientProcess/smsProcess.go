package clientProcess

import (
	"code/Project_practice/CommSystem_海量通讯系统/common/message"
	"code/Project_practice/CommSystem_海量通讯系统/common/utils"
	"encoding/json"
	"fmt"
)

// 发送群聊信息
func SendGroupMes() {
	for {
		var content string
		fmt.Print("发送：")
		fmt.Scanln(&content)

		// 进行判断，如果发送的消息为 “0”
		// 那么就退出群发功能
		if content == "0" {
			fmt.Println("======成功退出发送消息功能!======")
			return
		}

		// 否则就将信息序列化并发送给服务器
		smsMes := message.SmsMes{}
		smsMes.Content = content
		smsMes.AcoNum = CurUser.AcoNum


		data, err := json.Marshal(smsMes)
		if err != nil {
			fmt.Println("smsProcess json.Marshal(smsMes) err: ", err)
			return
		}

		// 创建 Message结构体实例
		mes := message.Message{
			Type: message.SmsMesType,
			Date: string(data),
		}
		// 序列化
		data, err = json.Marshal(mes)
		if err != nil {
			fmt.Println("smsProcess json.Marshal(mes) err: ", err)
			return
		}

		// 发送消息
		err = utils.WritePkg(CurUser.Conn, data)
		if err != nil {
			fmt.Println("smsProcess utils.WritePkg err: ", err)
			return
		}
	}
}

// 接收群聊消息
func ReceiveGroupMes(mes *message.Message) {
	// 反序列化接收到的消息 mes.Data
	var smsResMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Date), &smsResMes)
	if err != nil {
		fmt.Println("clientProcess json.Unmarshal([]byte(mes.Date), &smsResMes) err: ", err)
		return
	}
	info := fmt.Sprintf("\n接收：%v: %v\n", smsResMes.AcoNum, smsResMes.Content)
	fmt.Println(info)
}