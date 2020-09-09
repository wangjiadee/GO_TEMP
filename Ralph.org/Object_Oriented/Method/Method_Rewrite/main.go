package main

import (
	"fmt"

)

// Class
type Person struct {
	name string
	age  int
}

// Class method
func (p *Person) PrintInfo() {
	fmt.Println("父类中的方法")
}

type Student struct {
	Person
	score float64
}

func (p *Student) PrintInfo() {
	fmt.Println("子类的方法")
}

// =========================main function============================
func main() {

	// If the method name in the parent class is the same as the method name
	// in the subclass, then the method in the subclass is called through the
	// object of the subclass.   Method rewriting
	var stu Student
	stu.PrintInfo() 
	stu.Person.PrintInfo()
}

