package main

import "fmt"

func main() {
	s := []int{3, 5, 6, 7, 8, 9}
	//第一个值：截取的起始位置
	// 第二个值;截取的终止位置（不包含该值的）
	// 第三个值：用来计算容量，容量指的是切片中最多能够容纳多少元素。
	//容量=第三个值减去第一个值。
	//长度=第二个值减去第一个值
	//s1 := s[1:3:5]
	//fmt.Println(s1)
	//fmt.Println(cap(s1))
	//fmt.Println(len(s1))

	//s1 := s[:]
	//fmt.Println(s1)
	//fmt.Println(len(s1))
	//fmt.Println(cap(s1))

	s1 := s[1:]
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1)


	//s1:=s[:3]
	//	fmt.Println(s1)
	//	fmt.Println(len(s1))
	//	fmt.Println(cap(s1))

	//s1:=s[1:3]
	//fmt.Println(len(s1))
	//fmt.Println(cap(s1))
	//fmt.Println(s1)
}
