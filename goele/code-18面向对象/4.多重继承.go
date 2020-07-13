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


func main() {
	var stu Student
	//stu.age = 18
	fmt.Println(stu.Person.age)
	stu.id = 101
	//fmt.Println(stu.Person.Object.id)
	fmt.Println(stu.id)
}

