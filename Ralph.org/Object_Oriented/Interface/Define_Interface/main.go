package main

import "fmt"

// Define interface
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
	person.SayHello() // 调用的是Student方法实现的SayHello方法
	person = &tea
	person.SayHello()
}
