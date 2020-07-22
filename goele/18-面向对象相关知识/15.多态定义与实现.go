/*======================================================================
author：Ralph
time：2020-07-21
effect：go 多态相关
======================================================================*/
package main

import "fmt"

// 定义类
type Student struct {
}

type Teacher struct {
}

// 定义接口使用的方法
func (s *Student) SayHello() {
	fmt.Println("hello teacher!")
}

func (t *Teacher) SayHello() {
	fmt.Println("hello student!")
}

// 定义Personer的接口
type Personer interface {
	SayHello()
}

// 定义多态
func WhoSayHi(h Personer) {
	h.SayHello()
}

// =========================main function============================
func main() {
	var stu Student
	var teacher Teacher
	WhoSayHi(&stu)
	WhoSayHi(&teacher)
}
