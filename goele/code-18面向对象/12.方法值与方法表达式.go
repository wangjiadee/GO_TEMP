package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) PrintInfo() {
	fmt.Println(*p)
}

func main() {
	per := Person{"张三", 18}
	//per.PrintInfo()

	// 方法值。
	f := per.PrintInfo  //对象名。方法名字
	fmt.Printf("%T",f)
	f()
	// 方法表达式
	//f:=(*Person).PrintInfo // 要和结构体的类型一致的
	//f(&per)
}
