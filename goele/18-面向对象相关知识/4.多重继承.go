package main

import "fmt"

type Student struct {
	Person
	score float64
}
type Person struct {
	Object
	name string
	age  int
}
type Object struct {
	id int
}

// ***************************主函数**********************
func main() {
	var stu Student
	stu.age = 19
	fmt.Println(stu.Person.age)
	stu.id = 666
	fmt.Println(stu.id)
	fmt.Println(stu.Person.Object.id)
}
