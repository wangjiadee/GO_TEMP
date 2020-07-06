package main

import "fmt"

/*
func AddResult(num1 int,num2 int) int { // 表示指定函数返回的数据的类型。
	var sum int
	sum = num1+num2
	return sum //将变量sum中存储的值返回。
}
*/
// 表明：最终会返回整型变量sum中的值。
// 在函数体中没有必要在重新创建sum变量。
/*
func AddResult(num1 int,num2 int) (sum int) {
	// var sum int
	sum = num1+num2
	return sum //将变量sum中存储的值返回。
}
*/
func AddResult(num1 int,num2 int) (sum int) {
	// var sum int
	sum = num1+num2
	return  //如果已经指定了返回的变量的名称，那么return后面可以不用在加上变量的名称。
}
func GetResult()(num1 int,num2 int){
	//var num1 int = 10
	// var num2 int = 20
	num1=10
	num2=20
	//return num1,num2 //表明返回两个变量的值。
	return 
}
func main() {
	var s int
	s = AddResult(3,6) //将函数返回的结果赋值给了变量s.
	fmt.Println(s/3)

	var s1 int
	var s2 int
	s1,s2=GetResult() // 在这里，函数返回两个值，所以定义两个变量来进行接收。
	fmt.Println("s1=",s1)
	fmt.Println("s2=",s2)
}
