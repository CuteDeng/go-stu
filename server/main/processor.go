package main

import (
	"fmt"
	"go-stu/common/message"
	"go-stu/server/process"
	"go-stu/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) start() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器协程退出")
				return err
			} else {
				fmt.Println("readPkg err:", err)
				return err
			}
		}
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(&mes)
		if err != nil {
			return err
		}

	}
}
