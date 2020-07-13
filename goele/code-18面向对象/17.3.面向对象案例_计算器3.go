package main

/*
import "fmt"


// 加法类
type Add struct {
	Object
}

func (a *Add) GetResult() int { // 方法的实现要和接口中方法的声明保持一致
	return a.numA + a.numB
}

// 减法类
type Sub struct {
	Object
}

func (s *Sub) GetResult() int {
	return s.numA - s.numB
}

type Object struct {
	numA int
	numB int
}
type Resulter interface {
	GetResult() int
}

// 对象问题
// 1: 定义一个新的类
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
*/
func main() {
	//	var operator OperatorFactory
	//num := operator.CreateOperator("-", 20, 10)
	//fmt.Println(num)

}
