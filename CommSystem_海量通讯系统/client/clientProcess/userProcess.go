package clientProcess

import (
	"code/Project_practice/CommSystem_海量通讯系统/common/message"
	"code/Project_practice/CommSystem_海量通讯系统/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

// 处理客户端的登录功能的消息的接收和发送
func ClientLogin(conn net.Conn) {
	// 创建 用户的结构体实例
	logMes := message.User{}
	fmt.Print("请输入您的账号：")
	_, _ = fmt.Scanln(&logMes.AcoNum)
	fmt.Print("请输入您的密码：")
	_, _ = fmt.Scanln(&logMes.AcoPwd)
	fmt.Println("正在登录，请稍等...")

	// 获取输入的信息, 并将信息序列化
	logJsonRe, err := json.Marshal(&logMes)
	if err != nil {
		fmt.Println("json.Marshal(logMes) err: ", err)
		return
	}
	// 创建 message结构体, 并将发送的信息保存到实例对象中
	mes := message.Message{
		Type: message.LoginMesType,
		Date: string(logJsonRe),
	}

	// 将 mes实例对象序列化
	mesJsonRe, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err: ", err)
		return
	}

	// 4、发送消息的长度
	err = utils.WritePkg(conn, mesJsonRe)
	if err != nil {
		fmt.Println("client utils.WritePkg(conn, mesJsonRe) err: ", err)
	}

	// 5、接收对方的响应
	mesRes, err := utils.ReadPkg(conn)
	if err != nil {
		fmt.Println("ClientLogin utils.ReadPkg(conn) err: ", err)
		return
	}
	// 创建 登录返回信息的结构体实例
	var cLoginUJRes message.ResMes
	// 将接收到的信息反序列化
	err = json.Unmarshal([]byte(mesRes.Date), &cLoginUJRes)

	// 登录成功
	if cLoginUJRes.Code == 200 {
		// 在客户端保存 登录成功用户 的 账号 和 与客户端的连接
		CurUser.Conn = conn
		CurUser.AcoNum = logMes.AcoNum

		fmt.Printf("========用户：%v登录成功========\n", cLoginUJRes.Error)

		// 将在线用户保持下来
		for _, v := range cLoginUJRes.OnlineUserList {
			clientUserOnline = append(clientUserOnline, v)
		}

		// 持续和服务器进行连接
		go CommWithServer(conn)

		// 显示二级菜单
		for{
			View2()
		}
	// 登陆失败
	} else {
		fmt.Println(cLoginUJRes.Code, cLoginUJRes.Error)
		return
	}
}

// 处理客户端的注册功能
func Enroll(conn net.Conn) {
	// 创建注册结构的实例
	enroll := message.User{}

	// 输入信息
	fmt.Print("请输入您的账号：")
	_, _ = fmt.Scanln(&enroll.AcoNum)
	fmt.Print("请输入您的密码：")
	_, _ = fmt.Scanln(&enroll.AcoPwd)
	fmt.Print("请输入您的用户名：")
	_, _ = fmt.Scanln(&enroll.AcoName)
	fmt.Println("正在注册，请稍等...")

	// 获取输入的信息, 并将信息序列化
	enrJsonRe, err := json.Marshal(&enroll)
	if err != nil {
		fmt.Println("json.Marshal(logMes) err: ", err)
		return
	}
	// 创建 message结构体, 并将发送的信息保存到实例对象中
	mes := message.Message{
		Type: message.EnrollMesType,
		Date: string(enrJsonRe),
	}

	// 将 mes实例对象序列化
	mesJsonRe, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err: ", err)
		return
	}

	// 4、发送消息的长度
	err = utils.WritePkg(conn, mesJsonRe)
	if err != nil {
		fmt.Println("client utils.WritePkg(conn, mesJsonRe) err: ", err)
	}

	// 5、接收对方的响应
	mesRes, err := utils.ReadPkg(conn)
	if err != nil {
		fmt.Println("ClientLogin utils.ReadPkg(conn) err: ", err)
		return
	}
	// 创建 登录返回信息的结构体实例
	var cLoginUJRes message.ResMes
	// 将接收到的信息反序列化
	err = json.Unmarshal([]byte(mesRes.Date), &cLoginUJRes)

	// 输出状态码

	// 注册成功
	if cLoginUJRes.Code == 698 {
		fmt.Printf("成功注册！自动为您转到登陆界面...\n")
		fmt.Printf("======================================\n\n")
		ClientLogin(conn)

	} else if cLoginUJRes.Code == 929 {
		fmt.Println(cLoginUJRes.Code, cLoginUJRes.Error)
		return
	}
}