package main

import "fmt"

// 子类
type Student struct {
	*Person //指针匿名字段
	score   float64
}

// 父类
type Person struct {
	id   int
	name string
	age  int
}

// ***************************主函数**********************
func main() {
	var stu Student = Student{&Person{101, "Ralph", 19}, 200}
	fmt.Println(stu) //{0xc0000044a0 200}
	fmt.Println(stu.name)
	fmt.Println(stu.id)
}
