package main

import "fmt"

func main() {
	var Numbers [5]int = [5]int{1, 2, 3, 4, 5}
	getPrint(Numbers)
}
func getPrint(n [5]int) {
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
}

//在函数中修改原来的数组 是不会影响的值
func main() {
	var bun [5]int = [5]int{1, 2, 3, 4, 5}
	GetP(bun)
	for i := 0; i < len(bun); i++ {
		fmt.Println(bun[i])
	}
}
func GetP(n [5]int) {
	//for i:=0;i<len(n);i++{
	//	fmt.Println(n[i])
	//}
	n[2] = 10
}
