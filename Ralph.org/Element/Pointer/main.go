/*
 * @Author: your name
 * @Date: 2020-09-02 09:46:14
 * @LastEditTime: 2020-09-07 11:53:53
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \Workbenchc:\Users\80288284\go\src\github.com\Ralph.org\element\go_package_method\cala\tempCodeRunnerFile.go
 */
package main

import "fmt"

func Pointer() {
	num := 10
	PTest(num) //output 10 recover Ptest()
	fmt.Println(num)
}

func PTest(n int) {
	n = 30
	fmt.Println(n)
}

// Define pointer
func De_pointer() {
	a := 520
	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p) //根据内存地址取出值
	*p = 222
	fmt.Println("a=", a)
}

func Pointer_nil() {
	var p *int
	fmt.Println(p) // nil pointer

}

// Pointer as function parameter
func pointer_para() {
	num := 10
	// Update(num) Pointer parameter must &
	Update(&num)
	fmt.Println(num)
}
func Update(p *int) {
	*p = 60
}

// 数组指针 定义数组和指针 指针指向数组
func Arr_Pointer() {
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var P *[10]int
	P = &nums
	fmt.Println(*P)      //get pointer in arr
	fmt.Println((*P)[3]) // []的运算优先级高于*
	fmt.Println(P[0])
	// Traverse pointer
	for i := 0; i < len(P); i++ {
		fmt.Println(P[i])
	}
	UpdateArr(P)
	fmt.Println(nums)
}
func UpdateArr(P *[10]int) {
	P[0] = 100
}

// Pointer_Arr
// 一个数组里面都是指针
func Pointer_Arr() {
	var p [2]*int
	var i int = 10
	var j int = 20
	p[0] = &i
	p[1] = &j
	fmt.Println(p)     // 打印的结构都是地址 [0xc0000140b0 0xc0000140b8]
	fmt.Println(*p[0]) // 不要加括号
	fmt.Println(*p[1])
	//循环
	//for i:=0;i<len(p);i++{
	//	fmt.Println(*p[i])
	//}
	for k, value := range p {
		fmt.Println(k)
		fmt.Println(*value)
	}

}

func Pointer_Slice() {
	s := []int{1, 2, 3, 4, 5, 6}
	var p *[]int
	p = &s
	fmt.Println(*p)
	fmt.Println((*p)[0])
	//fmt.Println(p[0])     //错误*********

	// update
	(*p)[0] = 200
	fmt.Println(s)

	//for i:=0;i<len(*p) ; i++ {
	// fmt.Println((*p)[i])
	//}

	for k, value := range *p {
		fmt.Println("k=", k)
		fmt.Println("value=", value)
	}
}

// struct and Pointer
type Student struct {
	id   int
	name string
	age  int
}

func Str_Po() {
	stu := Student{101, "张三", 18}
	var p *Student
	p = &stu

	fmt.Println(*p)
	fmt.Println((*p).name)
	fmt.Println(p.name)
	//// 修改结构体中的值。
	p.age = 20
	fmt.Println(stu)

	UpdateStruct(p)
	fmt.Println(stu)

}
func UpdateStruct(stu *Student) {
	stu.age = 21
}

// Multi-level pointer
func Mul_Pointer() {
	var a int = 10
	var p *int    //一级指针
	p = &a
	fmt.Println(*p)
	var pp **int    //二级指针
	pp = &p
	**pp = 200
	fmt.Println(a)

}

//////////////////////// start main function
func main() {
	// Pointer()
	// De_pointer()
	// Pointer_nil()
	// pointer_para()
	// Arr_Pointer()
	// Pointer_Arr()
	// Pointer_Slice()
	// Str_Po()
	Mul_Pointer()
}

