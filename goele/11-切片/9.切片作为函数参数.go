package main

import "fmt"

func main() {
	//s := []int{1, 2, 3, 4, 5}
	s:=make([]int,10)
	Init(s)
	fmt.Println("s=",s)
}
func Init(num []int) {
	for i:=0;i<len(num) ;i++  {
		num[i]=i
	}
	/*
	for i := 0; i < len(num); i++ {
		fmt.Println(num[i])
	}
	*/
}
