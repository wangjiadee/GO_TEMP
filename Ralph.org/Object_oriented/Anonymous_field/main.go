package main

import "fmt"

// 定义一个Person类
type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型，没有名字，继承了Person的成员 匿名字段，那么默认Student就包含了Person的所有字段
	id     int
	addr   string
}

func main() {
	//顺序初始化
	var s1 Student = Student{Person{"mike", 'm', 18}, 1, "bj"}
	//自动推导类型
	s2 := Student{Person{"ricky", 'm', 2}, 1, "22"}
	//指定成员初始化，没有初始化的成员会自动赋值为0
	s3 := Student{id: 2}
	fmt.Println(s1, s2, s3)
}

