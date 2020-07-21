/*======================================================================
author：Ralph
time：2020-07-21
effect：go方法继承
======================================================================*/
package main

import (
	"fmt"
)

// 子类
type Student struct {
	Person
	score float64
}

// 父类
type Person struct {
	id   int
	name string
	age  int
}

// 父类的方法 也会有被子类继承
func (p *Person) PrintInfo() {
	fmt.Println(*p)
}

// =========================main function============================
func main() {
	stu := Student{Person{1000, "YY", 19}, 300}
	stu.PrintInfo()
}
