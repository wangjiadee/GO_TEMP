package main

import "fmt"

func main() {
	Test(11)
}
func Test(n int) {
	defer TestRecover()
	var num [10] int
	num[n] = 12
	fmt.Println(num)
	fmt.Println("aaaaa")

}
func TestRecover(){
	fmt.Println(recover())
}

