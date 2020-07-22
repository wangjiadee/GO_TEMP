/*======================================================================
author：Ralph
time：2020-07-21
effect：go 面向对象练习-计算器-加减子类+方法+接口
======================================================================*/
package main

import "fmt"

// 加减子类
type Add struct {
	Object
}
type Sub struct {
	Object
}

// 父类
type Object struct {
	numA int
	numB int
}

// 对于类的方法
func (a *Add) GetResult() int {
	return a.numA + a.numB
}

func (s *Sub) GetResult() int {
	return s.numA - s.numB
}

// 定义接口
type Resulter interface {
	GetResult() int
}

// =========================main function============================
func main() {
	add := Add{Object{500, 20}}
	num := add.GetResult()
	fmt.Println(num)
	sub := Sub{Object{1315, 1}}
	num1 := sub.GetResult()
	fmt.Println(num1)
}
