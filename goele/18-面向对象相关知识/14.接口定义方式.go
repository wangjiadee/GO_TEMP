/*======================================================================
author：Ralph
time：2020-07-21
effect：go 接口定义的方式
======================================================================*/
package main

import "fmt"

// 定义接口
type Personer interface {
	SayHello()
}

type Student struct {
}

func (s *Student) SayHello() {
	fmt.Println("fucking teacher!")
}

type Teacher struct {
}

func (t *Teacher) SayHello() {
	fmt.Println("fucking student")
}

// =========================main function============================
func main() {
	// 对象名.方法名字
	var stu Student
	stu.SayHello()
	var tea Teacher
	tea.SayHello()

	// 通过接口变量来调用,必须都实现接口中所声明的方法。
	var person Personer
	person = &stu
	person.SayHello() // 调用的是Student 实现的SayHello方法
	person = &tea
	person.SayHello()
}
