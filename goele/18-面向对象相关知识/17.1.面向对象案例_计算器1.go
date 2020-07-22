/*======================================================================
author：Ralph
time：2020-07-21
effect：go 面向对象练习-计算器-类+方法
======================================================================*/
package main

import "fmt"

// 定义一个类
type Object struct {
}

// 添加方法
func (o *Object) GetResult(num1 int, num2 int, op string) int {
	var result int
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	}
	return result
}

// =========================main function============================
func main() {
	var obj Object
	NUM := obj.GetResult(500, 20, "+")
	fmt.Println(NUM)
}
