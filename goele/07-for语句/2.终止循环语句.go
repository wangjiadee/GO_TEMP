package main

import "fmt"

// break 方法终止循环
func main() {
	var i int
	for i = 0; i < 10; i++ {
		if i == 3 {
			break
		}
	}
	fmt.Println(i)
}

// 用户登录系统模拟
func main() {
	var username string
	var password string
	var count int
	for {
		fmt.Println("Pelase input username:")
		fmt.Scan(&username)
		fmt.Println("Pelase input password:")
		fmt.Scan(&password)
		if username == "mdzz" && password == "sb" {
			fmt.Println("loading success!")
			break
		} else {
			count++
			if count >= 3 {
				fmt.Println("Error more than three")
				break
			}
			fmt.Println("pelase agian!")
		}
	}
}

// continue 退出本次循环
func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}


// 用continue实现：计算1到100(含)之间的除了能被7整除之外所有整数的和。
func main() {
	var sum int
	for i:= 1;i<=100;i++ {
		if i % 7 ==0{
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}