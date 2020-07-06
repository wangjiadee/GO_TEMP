package main

import "fmt"

func main()  {
   // 1: 接收考试成绩
   var score float64
   fmt.Println("请输入考试成绩:")
   fmt.Scan(&score)
   // 2: 对考试成绩进行判断
   if score >= 90 {
   		 fmt.Println("A")
   } else if score >= 80 {
   		fmt.Println("B")
   } else if score >= 70 {
   		fmt.Println("C")
   } else if score >= 60 {
   		fmt.Println("D")
   } else {
   		fmt.Println("E")
   }
   // 3: 打印相应的信息
}
