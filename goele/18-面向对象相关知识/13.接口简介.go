/*======================================================================
author：Ralph
time：2020-07-21
effect：go 接口介绍
======================================================================*/
package main

import "fmt"

// 加法类
type Add struct {
	Object
}

// 加法类的方法
func (p *Add) GetResult(a int, b int) int {
	p.num1 = a
	p.num2 = b
	return p.num1 + p.num2
}

// 减法类
type Sub struct {
	Object
}

// 减法类的方法
func (s *Sub) GetResult(a int, b int) int {
	s.num1 = a
	s.num2 = b
	return s.num1 - s.num2
}

// 减法加法类的父类
type Object struct {
	num1 int
	num2 int
}

// =========================main function============================
func main() {
	var sub Sub
	n := sub.GetResult(10, 8)
	fmt.Println(n)
	var add Add
	n2 := add.GetResult(20, 500)
	fmt.Println(n2)
}
