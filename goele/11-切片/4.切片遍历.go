package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	/*
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	*/
	for _,v:= range s{
		//fmt.Println("i=",i)
		fmt.Println("v=",v)
	}
}
