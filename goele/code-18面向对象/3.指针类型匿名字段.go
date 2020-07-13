package main

import "fmt"

type Student struct {
	*Person // 匿名字段
	score float64
}
type Person struct {
	id   int
	name string
	age  int
}

func main()  {
	var stu Student=Student{&Person{101,"张三",18},90}
	fmt.Println(stu) //{0xc0000044a0 90}
	fmt.Println(stu.name)
	fmt.Println(stu.id)

}
