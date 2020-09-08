package main

import (
	"fmt"
	"sort"

)

func homework() {
	sa := make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
	}
	fmt.Println(sa)
}

// sort
func sort_example() {
	var a [5]int = [5]int{1, 2, 4, 53, 34}
	sort.Ints(a[:])
	fmt.Println(a)
}

// ####################start main function

func main() {
	homework()
	sort_example()
}

