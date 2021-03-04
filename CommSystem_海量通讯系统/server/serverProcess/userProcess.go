package serverProcess

import (
	"code/Project_practice/CommSystem_海量通讯系统/common/message"
	"code/Project_practice/CommSystem_海量通讯系统/common/utils"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
)

// 处理服务器的登录功能
func ServerLogin(conn net.Conn, redisConn redis.Conn, mes *message.Message, login *message.User) (err error) {

	// 将mes信息反序列化
	err = json.Unmarshal([]byte(mes.Date), &login)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Date), loginUnJsonMes) err: ", err)
		return
	}
	// 声明一个 Message结构体实例
	var resMes = message.Message{}

	// 声明一个 LoginResMes结构体实例
	lgResMes := message.ResMes{}

	// 判断账号、密码是否正确
	res, err := redis.ByteSlices(redisConn.Do("HMGet", login.AcoNum, "pwd", "name"))
	if err == nil {
		// 密码正确
		if string(res[0]) == login.AcoPwd {
			lgResMes.Code = 200
			lgResMes.Error = string(res[1])
			// 将登录成功的用户 添加到 userOnline中
			AddUser(login.AcoNum, conn)

			// 将登录成功的用户的信息发送给当前在线 的用户
			NotifyOnlineUsers(login.AcoNum)

			// 遍历在线用户的 账号，并存放到 ResMes实例中
			for k, _ := range serverUserOnline {
				lgResMes.OnlineUserList = append(lgResMes.OnlineUserList, k)
			}
		// 密码错误
		} else {
			lgResMes.Code = 501
			lgResMes.Error = "账号或者密码错误！"
		}
	// 用户不存在
	} else if err == redis.ErrNil {
		lgResMes.Code = 500
		lgResMes.Error = "该用户不存在, 请注册！"
	} else {
		lgResMes.Code = 404
		lgResMes.Error = "位置错误"
	}

	// 将返回的信息序列化
	jLgResMes, err := json.Marshal(lgResMes)
	if err != nil {
		fmt.Println("json.Marshal(lgResMes) err: ", err)
		return
	}
	// 并把 类型和信息 保存到 resMes对象中
	resMes.Type = message.ResMesType
	resMes.Date = string(jLgResMes)

	// 再将 resMes序列化
	data, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) err: ", err)
		return
	}
	// 发送消息
	err = utils.WritePkg(conn, data)
	if err != nil {
		fmt.Println("utils.WritePkg(conn, data) err: ", err)
		return
	}
	return
}

// 处理服务器注册功能
func ServerEnroll(conn net.Conn, redisConn redis.Conn, mes *message.Message, enroll *message.User) (err error) {
	// 将mes信息反序列化
	err = json.Unmarshal([]byte(mes.Date), &enroll)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Date), loginUnJsonMes) err: ", err)
		return
	}
	// 声明一个 Message结构体实例
	var resMes = message.Message{}

	// 声明一个 LoginResMes结构体实例
	lgResMes := message.ResMes{}

	_, err = redis.String(redisConn.Do("HGet", enroll.AcoNum, "pwd"))

	// 这个账户存在
	if err == nil {
		lgResMes.Code = 929
		lgResMes.Error = "该用户已存在"
	} else {
		// 用户不存在，可以注册
		if err == redis.ErrNil {
			_, err = redisConn.Do("HSet", enroll.AcoNum, "pwd", enroll.AcoPwd, "name", enroll.AcoName, "safePwd", enroll.SafePwd)
			if err != nil {
				fmt.Println("redisConn.Do() HSet err: ", err)
				return
			}
			lgResMes.Code = 698
			lgResMes.Error = "注册成功"
		} else {
			fmt.Println("redis.String(redisConn.Do()) HGet err: ", err)
			return
		}
	}
	// 将返回的信息序列化
	jLgResMes, err := json.Marshal(lgResMes)
	if err != nil {
		fmt.Println("json.Marshal(lgResMes) err: ", err)
		return
	}
	// 并把 类型和信息 保存到 resMes对象中
	resMes.Type = message.ResMesType
	resMes.Date = string(jLgResMes)

	// 再将 resMes序列化
	data, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) err: ", err)
		return
	}
	// 发送消息
	err = utils.WritePkg(conn, data)
	if err != nil {
		fmt.Println("utils.WritePkg(conn, data) err: ", err)
		return
	}
	return
}

// 当一个用户上线时，这个用户将自己上线的信息发送给各个在线用户
func NotifyOnlineUsers(acoNum string) {
	for num, conn := range serverUserOnline {
		// 过滤自己
		if num == acoNum {
			continue
		}

		// 发送消息
		// 将自己的 账号和状态 放入UserStatus结构体中
		userStatus := message.UserStatus{
			AcoNum: acoNum,
			Status: message.Online,
		}
		// 序列化
		jsonUserStatus, err := json.Marshal(userStatus)
		if err != nil {
			fmt.Println("json.Marshal(userStatus) err: ", err)
			return
		}
		// 创建 Message 结构体
		mes := message.Message{
			Type: message.UserStatusType,
			Date: string(jsonUserStatus),
		}
		// 对 mes序列化
		data, err := json.Marshal(mes)
		if err != nil {
			fmt.Println("json.Marshal(mes) err: ", err)
			return
		}
		// 发送给在线用户列表中的各个用户
		err = utils.WritePkg(conn, data)
		if err != nil {
			fmt.Println("NotifyOnlineUsers WritePkg err: ", err)
			return
		}
	}
}

/*
	} else {
		fmt.Println("该账号已存在！")
		var temp int
		var safePwd string
		fmt.Print("请问您是否要修改密码(1.是，0.否)：")
		_, _ = fmt.Scanln(&temp)
		switch temp {
		case 0:
			return
		case 1:
			fmt.Print("请输入您的安全密码：")
			_, _ = fmt.Scanln(&safePwd)
			res, err := redis.String(conn.Do("HGet", u.AcoNum, "SafePwd"))
			if err != nil {
				fmt.Println("conn.Do() HMGet err: ", err)
				return
			}
			// 将获取到的安全密码与用户输入的相比较
			if safePwd == res {
				fmt.Print("安全密码正确！请重新输入您的密码：")
				_, _ = fmt.Scanln(&u.AcoPwd)
				_, err = conn.Do("HSet", u.AcoNum, "AcoPwd", u.AcoPwd)
				if err != nil {
					fmt.Println("conn.Do() HSet err: ", err)
					return
				} else {
					fmt.Println("修改成功！")
				}
			} else {
				fmt.Println("安全密码错误！")
			}
		}
	}
}
*/