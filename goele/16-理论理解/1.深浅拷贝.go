package main

import "fmt"

func main() {
	var num int = 10
	ap := &num
	Update(ap)
	fmt.Println(num)
}
func Update(p *int) {
	*p = 60
}
