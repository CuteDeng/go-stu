package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.start()
	if err != nil {
		fmt.Println("客户端和服务器协程错误：", err)
		return
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
