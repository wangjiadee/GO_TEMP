package main

import "fmt"

func main()  {
	var chinese int
	var math int
	fmt.Println("pelase input u chinese grade:")
	fmt.Scan(&chinese)
	fmt.Println("pelase input u math grade:")
	fmt.Scan(&math)
	if chinese > 70 && math == 100 {
		fmt.Println("奖励绿帽子！")
	}
}
