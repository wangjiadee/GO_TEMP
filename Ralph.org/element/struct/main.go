package main

import "fmt"

// Define struct
type Student struct {
	id   int
	name string
	age  int
	addr string
}

func Struct() {
	// struct init
	var s Student = Student{21, "ralph", 18, "sz"}
	fmt.Println(s)
	// Partial initialization
	var s1 Student = Student{name: "xiejer", age: 29}
	fmt.Println(s1)

	// single init
	var stur Student
	stur.name = "xiaosan"
	stur.addr = "dg"
	fmt.Println(stur)
}

// strcut & arr
func STRARR() {
	var arr [3]Student = [3]Student{
		Student{101, "张三", 18, "北京"},
		Student{102, "李四", 18, "北京"},
		Student{103, "王五", 19, "北京"},
	}
	fmt.Println(arr)

	// output fmt:
	fmt.Println(arr[0])
	fmt.Println(arr[0].age)
	// change value
	arr[0].age = 20
	fmt.Println(arr)

	// 通过循环来输出结构体数组中的内容。
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i].age)
	}

	for k, v := range arr {
		fmt.Println(k)
		fmt.Println(v.age)
	}
}

func Slice_struct() {
	var s []Student = []Student{
		Student{101, "张三", 18, "北京"},
		Student{102, "李四", 18, "北京"},
	}

	fmt.Println(s[0])
	fmt.Println(s[0].age)
	s[0].age = 20
	fmt.Println(s)

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

func MAPstruct() {
	m := make(map[int]Student)
	m[1] = Student{101, "张三", 18, "北京"}
	m[2] = Student{102, "李四", 18, "北京"}
	fmt.Println(m)
	fmt.Println(m[1])
	fmt.Println(m[1].name)

	delete(m, 2)
	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key)
		fmt.Println(value.age)
	}
}

func EXP_STRUCT() {
	stu := make([]Student, 3)
	InitData(stu)
	GetMax(stu)
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

// ###############start main function

func main() {
	// Struct()
	// STRARR()
	// Slice_struct()
	MAPstruct()
	EXP_STRUCT()
}

