package main

import "fmt"

func main()  {
	 var money float64
	 fmt.Println("请输入钱数")
	 fmt.Scan(&money)
	 // (2) 对钱数进行判断
	 if money > 2{
		 // (3) 如果钱数大于2元，条件成立，判断座位数。
		 // (3.1) 要求用户输入座位数.
		 var count int
		 fmt.Println("请输入座位数：")
		 fmt.Scan(&count)
		 // (3.2) 对输入的座位数进行判断。
		 if count > 0{
		 	fmt.Println("你可以坐下！！")
		 } else {
		 	fmt.Println("你只能站着！！")
		 }
	 } else {
		 // (4) 如果钱数不满足，只能输出”不能上车”
		 fmt.Println("你不能上车！！")
	 }



}
