package main

import (
	"fmt"

)

// sub class
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

// The methods of the futher class can be inherited by the Sub class
func (p *Person) PrintInfo() {
	fmt.Println(*p)
}

// =========================main function============================
func main() {
	stu := Student{Person{1000, "YY", 19}, 300}
	stu.PrintInfo()
}

