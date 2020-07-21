/*======================================================================
author：Ralph
time：2020-07-21
effect：go方法的重新
======================================================================*/
package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

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
	// 如果父类中的方法名称与子类中的方法名称一致，
	// 那么通过子类的对象调用的是子类中的方法。方法重写
	var stu Student
	stu.PrintInfo()
	stu.Person.PrintInfo()
}
