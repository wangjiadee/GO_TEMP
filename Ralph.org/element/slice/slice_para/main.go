package main

import "fmt"

func Slice_para() {
	s := make([]int, 5, 5)
	fmt.Println(s) //[0 0 0 0 0]
	Modify(s)
	fmt.Println(s) //[0 0 0 0 0]
}

func Modify(slice []int) {
	for i := 0; i < 5; i++ {
		slice = append(slice, i)
	}
	fmt.Println(slice) //[0 0 0 0 0 0 1 2 3 4]
}

/////////////////////start main function
func main() {
	Slice_para()
}

