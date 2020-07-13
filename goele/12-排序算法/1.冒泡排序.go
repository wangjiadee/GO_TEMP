package main

import "fmt"

func main() {
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
