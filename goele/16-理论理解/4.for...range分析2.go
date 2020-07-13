package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
		// 定义一个student 类型切片
		stu1 := Student{"MrWang", 100}
		stu2 := Student{"MrSun", 50}
		stu3 := Student{"MrQiu", 25}
		v := []Student{stu1, stu2, stu3}
		// 循环遍历 数组
		for _, value := range v {
			value.Age = value.Age + 1000 // 对每个人的年龄进行更新
		}
		// 打印切片的值
		fmt.Println(v)

}
