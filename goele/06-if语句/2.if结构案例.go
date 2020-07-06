package main

import "fmt"

func main() {
	var chinese int
	var math int
	fmt.Println("请输入语文成绩:")
	fmt.Scan(&chinese)
	fmt.Println("请输入数学成绩:")
	fmt.Scan(&math)
	if chinese > 70 && math == 100{
		fmt.Println("奖励100元")
	}
}