/*
 * @Author: your name
 * @Date: 2020-08-26 15:42:01
 * @LastEditTime: 2020-08-26 15:42:48
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \newgo\Ralph.org\element\go_syntax\if_syntax\main.go
 */
package main

import "fmt"

// eg about if syntax
func If_Eg() {
	fmt.Println("pelase input u age:")
	var age int
	fmt.Scan(&age)
	// if 的使用方式
	if age >= 18 {
		fmt.Println("u can buy Yellow vedio")
	}
}

// eg2 about if syntax
func If_Eg2() {
	var chinese int
	var math int
	fmt.Println("pelase input u chinese grade:")
	fmt.Scan(&chinese)
	fmt.Println("pelase input u math grade:")
	fmt.Scan(&math)
	if chinese > 70 && math == 100 {
		fmt.Println("奖励绿帽子！")
	}
}

// About if...else syntax
func Ifelse() {
	fmt.Println("pls input u age:")
	var age int
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("u can see yellow video!")
	} else {
		fmt.Println("studying hard now!")
	}
}

// ###########################start main function
func main() {
	// If_Eg()
	// If_Eg2()
	// Ifelse()
}
