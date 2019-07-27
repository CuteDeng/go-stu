package main

import "fmt"

var userId int
var userPwd string

func main() {
	var key string
	var loop = true
	for loop {
		fmt.Println("---------------------欢迎登陆聊天系统----------------------")
		fmt.Println("\t\t\t 1 登陆聊天室内")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）：")
		fmt.Scanf("%s\n", &key)
		fmt.Println(key)
		switch key {
		case "1":
			fmt.Println("登陆聊天室")
			loop = false
		case "2":
			fmt.Println("注册用户")
			loop = false
		case "3":
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

	if key == "1" {
		fmt.Println("请输入用户id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &userPwd)

		login(userId, userPwd)
		// if err != nil {
		// 	fmt.Println("登陆失败")
		// } else {

		// }

	} else if key == "2" {
		fmt.Println("进行用户注册")
	}

}
