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
	fmt.Println(*p)
}

// =========================main function============================
func main() {
	per := Person{"YAYA", 98}
	per.PrintInfo()  // input struct method

	// 方法值
	f := per.PrintInfo //对象名.方法名
	fmt.Printf("%T", f)
	f() //func(){YAYA 98}

	// 方法表达式
	f2 := (*Person).PrintInfo //和类 就是结构体一样
	f2(&per)                  //{YAYA 98}
}

