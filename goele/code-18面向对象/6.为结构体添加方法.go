package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
}

// 接收者
//func (s Student) PrintShow() {
//	fmt.Println(s)
//}
//func (s *Student) EditInfo() {
//	s.age = 20
//}
//func main() {
//	/*stu := Student{101, "张三", 18}
//	stu.PrintShow() // 完成对方法的调用, 将stu中的值，传递给了方法的s.
//	stu.EditInfo()
//	stu.PrintShow()*/
//}

//接受者
func (s Student) prinshow(){
	fmt.Println(s)
}


func (s1 *Student) editInfo() {
	s1.age	= 22
}
func main() {
	stu:= Student{
		id:   101,
		name: "xiejie",
		age:  20,
	}
	stu.prinshow()//將stu里面的值全部传递给了s
	stu.editInfo()
	stu.prinshow() //并没有对原成员值修改  修改成指针就可以了
}