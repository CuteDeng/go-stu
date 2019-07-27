package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-stu/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Print("读取客户端发送的数据...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		// err = errors.New("read pkg head err")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn read err:", err)
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		// err = errors.New("read pkg body err")
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.write err:", err)
		return
	}
	n, err = conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.write err:", err)
		return
	}
	return
}

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json unmarshal fail，err:", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在"
	}
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json marshal fail err:", err)
		return
	}

	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json marshal fail err:", err)
		return
	}

	err = writePkg(conn, data)
	return
}

func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器协程退出")
				return
			} else {
				fmt.Println("readPkg err:", err)
				return
			}
		}
		// fmt.Println("mes=", mes)
		err = serverProcessLogin(conn, &mes)
		if err != nil {
			return
		}

	}
}

func main() {
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")

	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	for {
		fmt.Println("等待客户端连接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err:", err)
		}
		go process(conn)
	}
}
