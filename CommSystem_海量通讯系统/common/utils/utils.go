package utils

import (
	"code/Project_practice/CommSystem_海量通讯系统/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

// 读取数据
func ReadPkg(conn net.Conn) (mes message.Message, err error) {
	// 接收数据的长度
	var buf = make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err == io.EOF {
		return
	} else if err != nil {
		fmt.Println("utils conn.Read(buf[:4]) err: ", err)
		return
	}

	// 接收数据本身
	// 1、将获取到的信息长度从 []byte类型 转换为 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	// 2、接收消息
	r, err := conn.Read(buf[:pkgLen])
	if r != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(buf[:pkgLen]) err: ", err)
		return
	}
	// 3、将消息反序列化
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen], mes) err: ", err)
		return
	}
	return
}

// 发送数据
func WritePkg(conn net.Conn, data []byte) (err error) {
	// 1、获得要发送的信息的长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:4], pkgLen)
	// 发送长度
	n, err := conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) err: ", err)
	}

	// 2、发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(mesJsonRe) err: ", err)
		return err
	}
	return
}