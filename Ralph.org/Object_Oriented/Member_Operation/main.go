package main

import "fmt"

// Sub class
type Student struct {
	Person
	score float64
}

// futher class
type Person struct {
	id   int
	name string
	age  int
}

// ///////////////////start main function
func main() {
	stu := Student{Person{101, "ralph", 18}, 90}
	stu1 := Student{Person{102, "mdzz", 13}, 110}
	fmt.Println(stu, stu1)
	stu.score = 1200
	fmt.Println(stu)
	fmt.Println(stu1.score)
	fmt.Println(stu1.Person.id)

}

