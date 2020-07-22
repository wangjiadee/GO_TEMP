/*======================================================================
author：Ralph
time：2020-07-21
effect：使用面向对象方式，实现两个数加减的计算器
======================================================================*/
package main

import "fmt"

type Add struct {
	Object
}

type Sub struct {
	Object
}

type Object struct {
	num1 int
	num2 int
}

// 定义的p 是同一个东西
func (p *Add) GetResult(a int, b int) int {
	p.num1 = a
	p.num2 = b
	return p.num1 + p.num2
}

func (p *Sub) GetResult(a int, b int) int {
	p.num1 = a
	p.num2 = b
	return p.num1 - p.num2
}

// =========================main function============================
func main() {
	var add Add
	var sub Sub
	n := add.GetResult(500, 20)
	N := sub.GetResult(1314, 1)
	fmt.Println(N)
}
