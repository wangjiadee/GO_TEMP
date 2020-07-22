/*======================================================================
author：Ralph
time：2020-07-21
effect：接口集成和转换
======================================================================*/
package main

import "fmt"

// 定义接口
type Humaner interface {
	SayHello()
}

type Personer interface {
	Humaner //继承父级接口
	Say()
}

// 定义方法
func (s *Student) SayHello() {
	fmt.Println("yayayayayaaya")
}

func (s *Student) Say() {
	fmt.Println("你好")
}

// 定义类
type Student struct {
}

// =========================main function============================
func main() {
	var stu Student
	var pre Personer
	pre = &stu
	pre.Say()
	// 子级接口可以调用父级接口
	pre.SayHello()

	var h Humaner
	h = pre
	h.SayHello()
}
