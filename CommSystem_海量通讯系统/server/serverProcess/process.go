package serverProcess

import (
	"code/Project_practice/CommSystem_海量通讯系统/common/message"
	"code/Project_practice/CommSystem_海量通讯系统/common/utils"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"net"
)

func Process(netConn net.Conn, pool *redis.Pool) {
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

		// 处理信息
		switch {
		// 处理登录信息
		case mes.Type == message.LoginMesType:
			// 创建登录结构体
			lm := message.User{}
			err = ServerLogin(netConn, redisConn, &mes, &lm)
			if err != nil {
				fmt.Println("serverProcess.ServerLogin(netConn, redisConn, &mes, &lm) err: ", err)
				return
			}

		// 处理注册信息
		case mes.Type == message.EnrollMesType:
			en := message.User{}
			err = ServerEnroll(netConn, redisConn, &mes, &en)
			if err != nil {
				fmt.Println("serverProcess.ServerEnroll(netConn, redisConn, &mes, &en) err: ", err)
				return
			}

		// 处理对方发送的信息
		case mes.Type == message.SmsMesType:
			err := ForwardMes(&mes)
			if err != nil {
			fmt.Println("serverProcess.ForwardMes(&mes) err: ", err)
			return
		}
		default:
			fmt.Println("输入的信息有误！")
		}
	}
}
