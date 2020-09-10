/*======================================================================
author：Ralph
time：2020-07-21
effect：go 面向对象练习-计算器-加减子类+方法中调用对象+接口
======================================================================*/
package main

import "fmt"

// 加减子类
type Add struct {
	Object
}
type Sub struct {
	Object
}

// 父类
type Object struct {
	numA int
	numB int
}

// 对于类的方法
func (a *Add) GetResult() int {
	return a.numA + a.numB
}

func (s *Sub) GetResult() int {
	return s.numA - s.numB
}

// 定义接口
type Resulter interface {
	GetResult() int
}

// 对象问题
type OperatorFactory struct {
}

// 2: 创建一个方法，在该方法中完成对象的创建

func (o *OperatorFactory) CreateOperator(op string, numA int, numB int) int {
	switch op {
	case "+":
		add := Add{Object{numA, numB}}
		return OperatorWho(&add)
	case "-":
		sub := Sub{Object{numA, numB}}
		return OperatorWho(&sub)

	default:
		return 0
	}
}

func OperatorWho(h Resulter) int {
	return h.GetResult()
}

// =========================main function============================
func main() {
	var operator OperatorFactory
	num := operator.CreateOperator("-", 20, 10)
	fmt.Println(num)
}

