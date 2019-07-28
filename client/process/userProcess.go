package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-stu/client/utils"
	"go-stu/common/message"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.dial err:", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal err:", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err:", err)
		return
	}
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.write err:", err)
		return
	}
	fmt.Printf("客户端，发送消息长度成功,长度为：%d,内容是：%s", len(data), string(data))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write err:", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err:", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
