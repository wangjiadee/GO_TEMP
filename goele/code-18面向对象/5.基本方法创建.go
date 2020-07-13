package main

import "fmt"

type Int int // type 为int类型指定了别名.
func (a Int) TestInt(b Int) Int {
	return a + b
}

func main() {
	var num Int = 10
	var num1 Int = 20
	n := num.TestInt(20)
	n1 := num1.TestInt(5)
	fmt.Println(n)
	fmt.Println(n1)

}
