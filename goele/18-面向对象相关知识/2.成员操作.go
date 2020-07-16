package main

import "fmt"

// 子类
type Student struct {
	Person
	score float64
}

// 父类
type Person struct {
	id   int
	name string
	age  int
}

// ***************************主函数**********************
func main() {
	var stu Student = Student{Person{101, "ralph", 18}, 90}
	var stu1 Student = Student{Person{102, "ralph", 18}, 90}
	fmt.Println(stu)
	// 单独修改成员的对象
	stu.score = 1000
	fmt.Println(stu)
	// 指定输出成员的对象
	fmt.Println(stu1.score)
	fmt.Println(stu1.Person.id)

}
