package main

import "fmt"

// Sub class (Pointer)
type Student struct {
	*Person
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
	stu := Student{&Person{101, "mdszz", 20}, 800}
	fmt.Println(stu) //输出Person 的内存地址
	fmt.Println(stu.name, stu.id)

}

