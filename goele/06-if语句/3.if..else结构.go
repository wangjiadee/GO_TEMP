package main

import "fmt"

func main()  {
	/*
	var age int = 15
	if age >= 18 {
		fmt.Println("恭喜你已经成年了!!")
	} else { // else 书写位置不要写错
		fmt.Println("你没有成年!!")
	}
	*/
	//小王的考试成绩大于等于90分,那么爸爸奖励他100元钱,否则

	//的话,爸爸就让小王跪方便面。

	// 1:定义一个变量存储小王的考试成绩。
	var score int
	// 2: 录入考试成绩
	fmt.Println("请输入考试成绩:")
	fmt.Scan(&score)
	// 3: 对考试成绩进行判断，然后打印出不同的结果。
	if score >= 90{
		fmt.Println("奖励100元")
	}else{
		fmt.Println("跪方便面")
	}
}
