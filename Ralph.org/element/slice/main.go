/*
 * @Author: Ralph
 * @Date: 2020-09-01 15:15:27
 * @LastEditTime: 2020-09-02 10:32:33
 * @LastEditors: Please set LastEditors
 * @Description: go的切片相关
 * @FilePath: \newgo\Ralph.org\element\slice\main.go
 */
package main

import "fmt"

func slice_struct() {
	var nums1 [5]int = [5]int{1, 2, 3, 4, 6}
	fmt.Println(nums1)
}

// use make to create slice
func make_slice() {
	s := make([]int, 3, 5)
	fmt.Println(s, len(s), cap(s))
}

// Slice traversal
func for_slice() {
	s := []int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func for_range_slice() {
	s := []int{12, 23, 34, 45, 56}
	for _, v := range s {
		fmt.Println(v)
	}
}

// Slice interception
func slice() {
	s := []int{3, 4, 5, 67, 78}
	s1 := s[1:3:5]
	fmt.Println(s1, len(s1), cap(s1))
	//第一个值：截取的起始位置
	//第二个值: 截取的终止位置（不包含该值的）
	//第三个值：用来计算容量，容量指的是切片中最多能够容纳多少元素。
	//cap 容量=第三个值减去第一个值。
	//len 长度=第二个值减去第一个值

	s2 := s[:3]
	fmt.Println(s2, len(s2), cap(s2))

}

// change slice value
func slice_change_l() {
	s := []int{1, 2, 3, 2, 4, 5, 6, 7, 8, 10}
	s1 := s[3:5]
	s[0] = 30
	s1[0] = 49
	fmt.Println(s, s1)
}

// append function app
func append_slice_l() {
	s := make([]int, 5, 8)
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4)
	fmt.Println(s, len(s), cap(s))
}

func copy_slice_l() {
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5, 6, 7}
	// s2把 s1 copy
	// copy(s1, s2)
	// fmt.Println(s1, s2)

	// s1 把 s2 copy
	copy(s2, s1)
	fmt.Println(s1, s2)

}

// Slice as function parameter
func slice_func_para_l() {
	// s := []int{123, 234, 345, 456, 567}
	s := make([]int, 10)
	Init(s)
	fmt.Println(s)
}

func Init(num []int) {
	for i := 0; i < len(num); i++ {
		num[i] = i
	}
}

/* for example slice
计算出一组整型数据之和。
1: 明确求的数据的个数。
2:输入要求和的数据，并且存储到切片中
3:求和运算
*/

func slice_example_l() {
	var count int
	fmt.Println("Pls input numbers count:")
	fmt.Scan(&count)
	s := make([]int, count)
	InitData(s)
	sum := SumAdd(s)
	fmt.Println(sum)

}

func SumAdd(num []int) int {
	var sum int
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}

func InitData(num []int) {
	for i := 0; i < len(num); i++ {
		fmt.Printf("Pls input  %d num!\n", i+1)
		fmt.Scan(&num[i])
	}
}

// ################start main function

func main() {
	// slice_struct()
	// make_slice()
	// for_slice()
	// for_range_slice()
	// slice()
	// slice_change_l()
	// append_slice_l()
	// copy_slice_l()
	// slice_func_para_l()
	slice_example_l()
}
