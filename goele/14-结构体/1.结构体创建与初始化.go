package main

import "fmt"

func main() {
	// 成员名称前面不能添加var.
	type Student struct {
		id   int
		name string
		age  int
		addr string
	}
	var s Student = Student{101, "张三", 18, "北京"}
	fmt.Println(s)
	// 部分初始化
	var s1 Student = Student{name: "李四", age: 18}
	fmt.Println(s1)

	var stu Student
	stu.id = 102
	stu.name = "老王"
	stu.age = 28
	stu.addr = "北京"
	fmt.Println(stu)
}
