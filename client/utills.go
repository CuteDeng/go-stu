package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-stu/common/message"
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
