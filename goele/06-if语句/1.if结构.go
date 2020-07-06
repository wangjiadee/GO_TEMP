package main

import "fmt"

func main()  {
	fmt.Println("pelase input u age:")
	var age int
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("u can buy Yellow vedio")
	}
}