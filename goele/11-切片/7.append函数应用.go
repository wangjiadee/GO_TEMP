package main

import "fmt"

func main() {
	s := make([]int, 5, 8)
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4)
	fmt.Println("len=", len(s))
	fmt.Println("cap=", cap(s))
	fmt.Println(s)
}
