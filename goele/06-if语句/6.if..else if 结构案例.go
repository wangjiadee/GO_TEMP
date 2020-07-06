package main

import "fmt"

func main()  {
	/*
	提示用户输入用户名，然后再提示输入密码，如果用户输入的用户名是“admin”并
	且密码是“88888”，则提示正确，否则，如果用户名是admin提示密码错误（用户名输入
	正确，密码输入错误）,如果密码是“88888”提示用户名错误（密码输入正确，用户名输入
	错误）。
	*/
	// 1: 接收用户输入的用户名和密码
	var userName string
	var userPwd string
	fmt.Println("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Println("请输入密码：")
	fmt.Scan(&userPwd)
	// 2: 对用户名和密码进行校验，如果用户名和密码都输入正确，给出"可以登录系统的提示"
	if userName == "admin" && userPwd == "88888" {
		fmt.Println("可以登录系统")
	} else  if userName == "admin" {
		fmt.Println("密码输入错误")
	} else if userPwd == "88888"  {
		fmt.Println("用户名错误！！")
	} else {
		fmt.Println("用户名和密码都输入错误！！")
	}
	// 3: 如果输入错误，要判断是否是用户名输入错误或者密码输入错误。
}
