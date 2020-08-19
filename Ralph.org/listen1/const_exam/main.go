/*
 * @Author: your name
 * @Date: 2020-08-19 14:57:56
 * @LastEditTime: 2020-08-19 14:58:18
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \newgo\Ralph.org\listen1\const_exam\main.go
 */
package main

import "fmt"

const (
	A = iota
	B
	C
	D = 8
	E
	F = iota
	G
)

func main() {
	fmt.Println(A, B, C, D, E, F, G)
	// 0 1 2 8 8 5 6 E没有写 和签名默认 iota 表示行数
}
