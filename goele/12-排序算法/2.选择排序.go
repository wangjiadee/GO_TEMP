package main

import "fmt"

func main() {
	// 考虑第一趟的情况
	// 1；找出切片中最小的数据。
	// 2: 和切片中第一个数进行位置的交换。
	s := []int{5, 9, 0, 2, 7}
	for j:=0;j<len(s)-1;j++{
		min := s[j]
		minIndex := j
		for i:=j+1;i<len(s);i++{
			if min>s[i]{
				min=s[i]
				minIndex=i
			}
		}
		if minIndex!=j{
			s[j],s[minIndex]=s[minIndex],s[j]
		}
	}
	fmt.Println(s)

/*
	// 第二趟
	min = s[1]
	minIndex = 1
	for i:=1+1;i<len(s);i++{
		if min>s[i]{
			min=s[i]
			minIndex=i
		}
	}
	if minIndex!=1{
		s[1],s[minIndex]=s[minIndex],s[1]
	}
	fmt.Println(s)
*/
}
