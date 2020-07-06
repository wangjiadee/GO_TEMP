package main

import "fmt"

func main()  {
	// 请用户输年份,再输入月份,输出该月的天数.(需要考虑闰年)
	/*
	闰年的判定(符合下面两个条件之一)
	年份能够被400整除.(2000)
	年份能够被4整除但不能被100整除.(2008)
	*/
	// 1: 定义变量存储月份与年
	var year int
	var month int
	var day int
	fmt.Println("请输入年：")
	fmt.Scan(&year)
	fmt.Println("请输入月：")
	fmt.Scan(&month)
	// 2: 判断输入的月份是否正确.
	if month >= 1 && month <= 12 {
		// 3: 如果 1，3,5,7,8,10,12  输入的天数31天
		switch month {
		case 1:
			fallthrough
		case 3:
			fallthrough
		case 5:
			fallthrough
		case 7:
			fallthrough
		case 8:
			fallthrough
		case 10:
			fallthrough
		case 12:
			day = 31
			// 4: 2 判断闰年 29天，28天
		case 2 :
			if year % 400 == 0 || year % 4 == 0 && year % 100 != 0 {
				day = 29
			} else {
				day = 28
			}
		default:	// 5: 其它月份输出的是30天
			day = 30
		}
		fmt.Println("天数：", day)
	} else {
		fmt.Println("输入月份错误！！")
	}
}
