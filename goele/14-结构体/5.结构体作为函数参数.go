package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	addr string
}

func main() {
	// 1: 输入学生信息
	stu := make([]Student, 3)
	InitData(stu)
	// 2: 进行比较
	GetMax(stu)
	// 3:打印结果

	//stu := Student{101, "张三", 18, "北京"}
	//PrintDemo(stu)
	//fmt.Println(stu)
}
func InitData(stu []Student) {
	for i := 0; i < len(stu); i++ {
		fmt.Printf("请输入第%d个学生的详细信息\n", i+1)
		fmt.Scan(&stu[i].id, &stu[i].name, &stu[i].age, &stu[i].addr)
	}
}
func GetMax(stu []Student) {
	var max int = stu[0].age
	var maxIndex int // 记录最大年龄学生信息在整个切片中下标。
	for i := 0; i < len(stu); i++ {
		if stu[i].age > max {
			max = stu[i].age
			maxIndex = i
		}
	}
	fmt.Println(stu[maxIndex])
}

//func PrintDemo(stu Student) {
//	stu.age = 20
//	// fmt.Println(stu)
//}
