package main

import (
	"code/Project_practice/CommSystem_海量通讯系统/server/model"
	"code/Project_practice/CommSystem_海量通讯系统/server/serverProcess"
	"fmt"
	"net"
	"time"
)

/*func process(netConn net.Conn, pool *redis.Pool) {
	// 使用线程池连接 redis
	redisConn := pool.Get()

	// 延时关闭连接
	defer netConn.Close()
	defer redisConn.Close()

	for {
		// 接收消息
		mes, err := utils.ReadPkg(netConn)
		if err == io.EOF {
			fmt.Printf("%v与服务器断开连接!\n", netConn.RemoteAddr().String())
			return
		} else if err != nil {
			fmt.Println("server utils.ReadPkg(conn) err: ", err)
			return
		}
		fmt.Println(mes)

		// 处理信息
		switch {
		// 处理登录信息
		case mes.Type == message.LoginMesType:
			// 创建登录结构体
			lm := message.User{}
			err = serverProcess.ServerLogin(netConn, redisConn, &mes, &lm)
			if err != nil {
				fmt.Println("serverProcess.ServerLogin(netConn, redisConn, &mes, &lm) err: ", err)
				return
			}

		// 处理注册信息
		case mes.Type == message.EnrollMesType:
			en := message.User{}
			err = serverProcess.ServerEnroll(netConn, redisConn, &mes, &en)
			if err != nil {
				fmt.Println("serverProcess.ServerEnroll(netConn, redisConn, &mes, &en) err: ", err)
				return
			}


		default:
			fmt.Println("输入的信息有误！")
		}
	}
}*/

func main() {
	pool := model.PoolInit("localhost:6379", 16, 0, 300 * time.Second)

	fmt.Println("服务器开始监听！")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen() err: ", err)
		return
	}
	// 延时关闭监听
	defer listen.Close()
	defer pool.Close()

	for {
		netConn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err: ", err)
			continue
		}
		// RemoteAddr() 可以返回远端网络地址
		fmt.Printf("%v成功连接到服务器!\n", netConn.RemoteAddr().String())

		// 开启一个协程
		go serverProcess.Process(netConn, pool)
	}
}
