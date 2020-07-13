package main

import "fmt"

func main() {
	/*
	var s []int
	s = append(s, 1, 2, 3, 4, 5, 89, 90)
	s[3] = 88 // 通过这种方式可以完成某个值的修改。
	fmt.Println(s[3])
	fmt.Println(s)
	*/
	/*
	s := []int{8, 9, 7, 10, 12, 13}
	s = append(s, 99, 100)
	s[0] = 78
	fmt.Println(s[0])
	fmt.Println(s)
	*/
	s := make([]int, 3, 10)
	for i:=0;i<len(s) ;i++  {
		s[i]=i+1
	}
	s=append(s,80)
	s[3]=90
	/*
	s[0]=10
	s[1]=20
	s[2]=30
	*/

	// s[3]=78 // 会造成下标越界。
	//s[4]=80
	fmt.Println(s)
}
