package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// s[0] = 30
	s1 := s[2:5]
	//fmt.Println(s)
	s1[0] = 80
	fmt.Println("s1=", s1)
	fmt.Println("s=",s)
}
