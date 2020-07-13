package main

import "fmt"

// 使用面向对象方式，实现两个数加减的计算器
// 1: 抽取对应的类
// 2: 添加相应的方法。
// 加法类
type Add struct {
	Object
}

func (p *Add) GetResult(a int, b int) int {
	p.num1 = a
	p.num2 = b
	return p.num1 + p.num2
}

// 减法类
type Sub struct {
	Object
}

func (p *Sub) GetResult(a int, b int) int {
	p.num1 = a
	p.num2 = b
	return p.num1 - p.num2
}

type Object struct {
	num1 int
	num2 int
}

func main() {
	/*var add Add
	n := add.GetResult(10, 5)
	fmt.Println(n)*/
	var sub Sub
	n := sub.GetResult(10, 5)
	fmt.Println(n)

}
