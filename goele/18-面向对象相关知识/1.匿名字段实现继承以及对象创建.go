package main

import "fmt"

type Student struct {
	Person // 匿名字段 只有类型 没有成员名字
	score  float64
}

type Teacher struct {
	Person
	salary float64
}

type Person struct {
	id   int
	name string
	age  int
}

// ***************************主函数**********************
func main() {
	var stu Student = Student{101, "ralph", 18}
	var stu1 Student = Student{102,"Tom",18}
	fmt.Println(stu)
	fmt.Println(stu1)

	全部初始化
	var stu Student=Student{Person{101,"张三",18},90}
	fmt.Println(stu)	

	部分初始化：
	var stu Student = Student{Person: {id: 100}}}
	fmt.Println(stu)
}
