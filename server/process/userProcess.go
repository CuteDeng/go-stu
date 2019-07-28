package process2

import (
	"encoding/json"
	"fmt"
	"go-stu/common/message"
	"go-stu/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
