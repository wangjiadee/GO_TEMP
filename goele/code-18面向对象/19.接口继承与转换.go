package main

import "fmt"

type Humaner interface {
	SayHello()
}
type Personer interface {
	Humaner
	Say()
}
type Student struct {
}

func (s *Student) SayHello() {
	fmt.Println("大家好")
}
func (s *Student) Say() {
	fmt.Println("你好")
}
func main() {
	var stu Student
	var per Personer
	per = &stu
	//per.Say()
	//per.SayHello() // 可以调用所继承的接口中的方法。

	  var h Humaner
	 h=per
	 // per=h
	  h.SayHello()


}
