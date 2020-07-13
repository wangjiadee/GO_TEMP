package main

import "fmt"

func main() {
	var i interface{}
	i = 123
	i = "abc"
	fmt.Println(i)
	var s []interface{}
	s = append(s, 123, "abc", 12.3)
	for j := 0; j < len(s); j++ {
		fmt.Println(s[j])
	}
}
