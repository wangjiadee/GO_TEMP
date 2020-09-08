package main

import (
	"fmt"

)

// Define Student struct
type Student struct {
	id   int
	name string
	age  int
}

// Define Teacher struct
type Teacher struct {
	id   int
	name string
}

// Method of receiving struct
func (s *Student) PrinShow() {
	fmt.Println(s)
}

func (t *Teacher) PrinShow() {
	fmt.Println(t)
}

// =========================main function============================
func main() {

	// If the receiver type is different,
	// even if the method name is the same,
	// it will be a different method.
	stu := Student{802, "yaya", 14}
	// The results of the two outputs are the same
	(&stu).PrinShow()
	stu.PrinShow()
	teacher := Teacher{01, "校长"}
	(&teacher).PrinShow()
	teacher.PrinShow()
}
