package main

import "fmt"

// if else 大致结构
func main() {
	var age int
	if age >= 18 {
		fmt.Println("u can see yellow video!")
	} else {
		fmt.Println("studying hard now!")
	}
}

// 自动交互加if else 语句
func main() {
	var score int
	fmt.Println("pelase input fucking score:")
	fmt.Scan(&score)
	if score >= 90 {
		fmt.Println("give u ￥100")
	} else {
		fmt.Println("fucking u !")
	}
}
