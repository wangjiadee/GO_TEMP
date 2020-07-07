package main

import "fmt"

// if else 结构案例
func main() {
	var money float64
	fmt.Println("Pelase input money:")
	fmt.Scan(&money)
	if money >= 2 {
		var count int
		fmt.Println("Pelase input count imformation:")
		fmt.Scan(&count)
		if count > 0 {
			fmt.Println("down!")
		} else {
			fmt.Println("up!")
		}
	} else {
		fmt.Println("gun !")
	}
}
