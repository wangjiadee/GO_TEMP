package main

import (
        "fmt"
        "sort"
	"flag"

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

//secret flag
func ParseArgs() {
	var (
		length int
		charset string
	)
	flag.IntVar(&length, "l",16,"-l create password len")
	flag.IntVar(&charset, "t","num","-l create password len")
}


// ####################start main function

func main() {
        homework()
        sort_example()
}

