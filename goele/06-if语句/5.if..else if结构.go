package main

import "fmt"

// if else if 基础语法结构
func main() {
	var score float64
	fmt.Println("input u score:")
	fmt.Scan(&score)
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}

// if else if 案例
func main() {
	var username string
	var password string
	fmt.Println("Pelase input u username:")
	fmt.Scan(&username)
	fmt.Println("Pelase input u Password:")
	fmt.Scan(&password)
	if username == "mdzz" && password == "sb" {
		fmt.Println("loading success!")
	} else if username == "mdzz" {
		fmt.Println("password Error!")
	} else if password == "sd" {
		fmt.Println("Username Error!")
	} else {
		fmt.Println("password or Username Error!")
	}
}
