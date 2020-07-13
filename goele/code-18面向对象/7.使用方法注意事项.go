package main

import "fmt"

type Student struct {
	id int
	name string
	age int
}
type Teacher struct {
	id int
	name string
}
func(s *Student) Show(){
	fmt.Println(s)
}
func(t *Teacher) Show(){
	fmt.Println(t)
}
func main() {
	// 如果接收者类型不同，即使方法的名字是相同的也是不同的方法。
  stu:=Student{101,"李四",18}
	//(&stu).Show()
	stu.Show()
  teacher:=Teacher{102,"老王"}
  teacher.Show()
}
