/*
 * @Author: Ralph
 * @Date: 2020-09-01 15:04:07
 * @LastEditTime: 2020-09-03 09:57:12
 * @LastEditors: Please set LastEditors
 * @Description: GO 三种排序算法
 * @FilePath: \newgo\Ralph.org\element\algorithm\main.go
 */
package main

import "fmt"

func MaoPao_sort() {
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var temp int
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				temp = s[j]
				s[j] = s[j+1]
				s[j+1] = temp
			}
		}
	}
	fmt.Println(s)

}

func Choose_sort() {
	// 考虑第一趟的情况
	// 1；找出切片中最小的数据。
	// 2: 和切片中第一个数进行位置的交换。
	s := []int{5, 9, 0, 2, 7}
	for j := 0; j < len(s)-1; j++ {
		min := s[j]
		minIndex := j
		for i := j + 1; i < len(s); i++ {
			if min > s[i] {
				min = s[i]
				minIndex = i
			}
		}
		if minIndex != j {
			s[j], s[minIndex] = s[minIndex], s[j]
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

func Insert_sort() {
	var arr [8]int = [8]int{90, 8, 34, 23, 56, 7, 1, 28}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	fmt.Print(arr)
}

// ############start main function

func main() {
	MaoPao_sort()
	Choose_sort()
	Insert_sort()
}
