/*
程序员自己不会调用该函数，但是如果程序员自己写的程序中出现了比较严重的异常，那么系统内部会调用该函数，终止整个程序的执行。
*/

package main

import "fmt"

func Test(n int) {
	var num [10]int
	num[n] = 12
	fmt.Println("Fucking!")
	// 引发异常，强制终止整个程序的执行
	panic("Error!")
	fmt.Println("Fucking too!")
}

// ===================================Main Function start==============================
func main() {
	Test(11)
}
