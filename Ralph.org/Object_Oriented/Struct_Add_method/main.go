package main

import "fmt"

// Add method for Struct

// Define Struct
type Student struct {
	id   int
	name string
	age  int
}

// Receive
func (s Student) PrintShow() {
	fmt.Println(s)
}

func (s1 *Student) EditInfo() {
	s1.age = 22
}

// ///////////////////start main function
func main() {
	stu := Student{
		id:   101,
		name: "mdzz",
		age:  20,
	}
	// Pass all the values in stu to s
	stu.PrintShow()
	stu.EditInfo()
	// It does not modify the original member value,
	// just modify it to a pointer.
	stu.PrintShow()
}

