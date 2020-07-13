package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
}

func main() {
	stu := Student{101, "张三", 18}
	var p *Student
	p = &stu

	//fmt.Println(*p)
	//fmt.Println((*p).name)
	//fmt.Println(p.name)
	//// 修改结构体中的值。
	//p.age = 20
	//fmt.Println(stu)
	UpdateStruct(p)
	fmt.Println(stu)

}
func UpdateStruct(stu *Student) {
	stu.age = 21
}
