package main

import "fmt"

func main() {
	// 下标从0开始计算
	var number [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(number[3])
}

func main() {
	// 允许部分赋值
	Numbers := [5]int{1, 2}
	fmt.Println(Numbers[4])
}

func main() {
	// 运行指定初始化某一个元素
	// 如果没有初始化就是0
	Numbers := [5]int{2: 5, 3: 6}
	fmt.Println(Numbers[2])
	fmt.Println(Numbers[1])
}

func main() {
	// 可以通过初始化自定义数组的长度
	Numbers := [...]int{7, 8, 5}
	fmt.Println(len(Numbers))
	// fmt.Println(Numbers[0])
}

func main() {
	var Numbers [5]int
	/*
		Numbers[0]=1
		Numbers[1]=2
		fmt.Println(Numbers[3])
	*/
	for i := 0; i < len(Numbers); i++ {
		Numbers[i] = i + 1
	}
	fmt.Println(Numbers[0])
}
