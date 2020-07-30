/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2020-04-27 21:59:02
 * @LastEditTime: 2020-07-30 21:31:28
 * @FilePath: \go\goele\20-异常处理\2.panic函数.go
 * @Effect: panic函数异常处理
 */

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
