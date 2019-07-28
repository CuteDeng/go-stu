package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-stu/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Print("读取客户端发送的数据...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn read err:", err)
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.write err:", err)
		return
	}
	n, err = this.Conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.write err:", err)
		return
	}
	return
}
