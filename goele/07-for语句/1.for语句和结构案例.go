package main

import (
	"fmt"
)

// for 语法结构
func main() {
	var i int
	for i = 1; i <= 10; i++ {
		fmt.Println("love bb!")
	}
	fmt.Println(i)
}

// for 循环案例1 求1-100的所有数字
func main() {
	var i int
	var sum int
	for i = 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}

// for 循环案例2 求1-100的偶数和
func main() {
	var sum int
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
