package main

import "fmt"

func main()  {
	/*
	李四的年终工作评定,如果定为A级,则工资涨500元,如果定为B级,则工资涨200元,
	如果定为C级,工资不变,如果定为D级工资降200元,如果定为E级工资降500元.
	设李四的原工资为5000,请用户输入李四的评级,然后显示李四来年的工资.
	*/
	// 1: 定义一个变量接收输入的评级信息。
	 var level string
	 var salary int = 5000
	 var b bool = true
	 fmt.Println("请输入评级信息:")
	 fmt.Scan(&level)
	// 2: 根据输入的评级信息，进行判断，从而决定工资是加还是减。
	/*
	 if level == "A" {
	 	salary += 500
	 } else if level == "B" {
	 	salary += 200
	 } else if level == "C" {

	 } else if level == "D" {
		salary -= 200
	 } else if level == "E" {
		 salary -= 500
	 } else {
	 	b = false
	 	fmt.Println("输入评级信息错误！！")
	 }
	*/

	// 使用switch 方式来解决该问题。
	switch level {
	case "A":
		salary += 500
	case "B":
		salary += 200
	case "C":
	case "D":
		salary -= 200
	case "E":
		salary -= 500
	default:
		b = false
		fmt.Println("输入评级信息错误！！")
	}
	// 3: 打印相应的工资。
	if b {
		fmt.Println("工资是：",salary)
	}
}
