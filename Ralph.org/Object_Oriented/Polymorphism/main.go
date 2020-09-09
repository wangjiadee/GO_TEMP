package main

import "fmt"

// Define class
type Student struct {
}

type Teacher struct {
}

// Define method
func (s *Student) SayHello() {
	fmt.Println("Fucking teacher!")
}

func (t *Teacher) SayHello() {
	fmt.Println("Fucking student!")
}

// Define interface
type Personer interface {
	SayHello()
}

// Define Polymorphism
func WhoSayHi(h Personer) {
	h.SayHello()
}

// =========================main function============================
func main() {
	var stu Student
	var teacher Teacher
	WhoSayHi(&stu)
	WhoSayHi(&teacher)
}

