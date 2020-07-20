package main

import "fmt"

// 指定一个方法叫做Int 类型为int
type Int int

func (a Int) TestInt(b Int) Int {
	return a + b
}

// =========================main function============================
func main() {
	var num Int = 20
	var num1 Int = 100
	n := num.TestInt(20)
	n1 := num1.TestInt(30)
	fmt.Println(n)
	fmt.Println(n1)
}
