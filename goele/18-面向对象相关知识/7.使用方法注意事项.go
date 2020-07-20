package main

import (
	"fmt"
)

// 定义结构体1
type Student struct {
	id   int
	name string
	age  int
}

// 定义结构体2
type Teacher struct {
	id   int
	name string
}

// 2个接受者  为结构体定义方法
func (s *Student) PrinShow() {
	fmt.Println(s)
}

func (t *Teacher) PrinShow() {
	fmt.Println(t)
}

// =========================main function============================
func main() {
	// 如果接收者类型不同，即使方法的名字是相同的也是不同的方法。
	stu := Student{802, "yaya", 14}
	stu.PrinShow() // 此用法和下面的用法是一样的
	// (&stu).PrinShow()
	teacher := Teacher{01, "校长"}
	(&teacher).PrinShow()
}
