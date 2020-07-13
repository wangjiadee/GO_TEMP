package main

import "fmt"

func main()  {
	v := [3]int{1, 2, 3}
	for i,value := range (v) {  //对形参进行遍历
		v[i]=3
		// 输出打印遍历值
		fmt.Println("value:",value) // 1 2 3
	}
	// 打印切片的值
	fmt.Println("v",v) //v [3 3 3]

}


