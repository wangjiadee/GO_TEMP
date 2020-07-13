package main

import "fmt"

func main() {
	Test(11)

}
func Test(n int) {
	var num [10] int
	num[n] = 12
	fmt.Println("hhh")
	// 程序员自己不会调用该函数，但是如果程序员自己写的程序中出现了比较严重的异常，那么系统内部会调用该函数，终止整个程序的执行。
	//panic("abc") // 引发异常，从而强制终止整个程序的执行。
	//fmt.Println("hello")
}
