package clientProcess

import (
	"fmt"
	"net"
	"os"
)

func View1(conn net.Conn) {
	// 1、打印菜单
	var temp int
	for {
		fmt.Print(`
------------Welcome communication system------------
					1、注  册
					2、登  录
					3、退  出
				请输入您要选择的功能：`)
		_, _ = fmt.Scanln(&temp)
		switch temp {
		// 注册
		case 1:
			Enroll(conn)

		// 登录
		case 2:
			ClientLogin(conn)
		// 退出
		case 3:
			fmt.Println("退出客户端！")
			return

		default:
			fmt.Println("输入有误！")
			return
		}
	}
}

func View2() {
	var temp int
	fmt.Print(`
	1、显示在线用户列表
		2、发送消息
		3、信息列表
		4、退出系统
		请选择(1-4):`)
	fmt.Scanln(&temp)

	switch temp {
	case 1:
		ShowOnlineUser()
	case 2:
		SendGroupMes()

	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入错误，请重新输入！")
	}
}

