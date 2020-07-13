package main

import "fmt"

func main() {
	var Numbers [5] int = [5]int{1, 2, 3, 4, 5}
	// 使用for循环方式遍历
	/*
	for i := 0; i < len(Numbers); i++ {
		fmt.Println(Numbers[i])
	}
	*/
	// 使用for range 方式遍历
	for i, v := range Numbers {
		fmt.Println("下标：", i)
		fmt.Println("值：", v)
	}
}

// 使用匿名变量
func main(){
	var numbers[5]int = [5]int{1,2,3,4,5}
	//for i:=0;i<len(numbers) ;i++  {
	//	fmt.Println(numbers[i])
	//}
	// _ 表示匿名变量
	for _,v:=range numbers{
		//fmt.Println("biao",i)
		fmt.Println("zhi",v)
	}
}