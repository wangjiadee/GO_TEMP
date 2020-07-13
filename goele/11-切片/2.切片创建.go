package main

import "fmt"

func main() {
	// var s[]int
	// s := []int{}
	s:=make([]int,3,5) // 长度不能大于容量。
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
