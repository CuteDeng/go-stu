package main

import (
	"fmt"
	"go-stu/client/process"
)

var userId int
var userPwd string

func main() {
	var key string
	for {
		fmt.Println("---------------------欢迎登陆聊天系统----------------------")
		fmt.Println("\t\t\t 1 登陆聊天室内")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）：")
		fmt.Scanf("%s\n", &key)

		switch key {
		case "1":
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case "2":
			fmt.Println("注册用户")
		case "3":
			fmt.Println("退出系统")
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

}
