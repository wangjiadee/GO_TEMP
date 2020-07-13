package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	addr string
}

func main() {

	var s []Student = []Student{
		Student{101, "张三", 18, "北京"},
		Student{102, "李四", 18, "北京"},
	}
	/*
		fmt.Println(s[0])
		fmt.Println(s[0].age)
		s[0].age = 20
		fmt.Println(s)
	*/

	// 循环遍历
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i].name)
	}
	for k, v := range s {
		fmt.Println(k)
		fmt.Println(v.id)

	}

	// 追加数据
	s = append(s, Student{103, "王五", 20, "北京"})
	fmt.Println(s)
}
