package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	addr string
}


func main() {
	var arr [3]Student = [3]Student{
		Student{101, "张三", 18, "北京"},
		Student{102, "李四", 18, "北京"},
		Student{103, "王五", 19, "北京"},
	}

	fmt.Println(arr)
	/*
	fmt.Println(arr[0])
	fmt.Println(arr[0].age)
	arr[0].age=20
	fmt.Println(arr)

	// 通过循环来输出结构体数组中的内容。
	for i:=0;i<len(arr) ;i++  {
		fmt.Println(arr[i].age)
	}
	*/

	//for k, v := range arr {
	//	fmt.Println(k)
	//	fmt.Println(v.age)
	//}

}
