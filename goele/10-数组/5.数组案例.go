package main

import "fmt"

func main() {
	// 从一个整数数组中取出最大的整数,最小整数,并且求总和,求平均值。
	// 1:定义数组
	var nums [5]int = [5]int{3, 6, 7, 8, 9}
	// 2:定义两个变量存储最大的值和最小的值。
	var max int = nums[0]
	var min int = nums[0]
	var sum int
	// 3: 循数组和最大值与最小值进行比较。
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
		sum = sum + nums[i]
	}
	fmt.Println("最大值为", max)
	fmt.Println("最小值为", min)
	fmt.Println("和为", sum)
	fmt.Printf("平均值为；%.2f", float64(sum)/5)

}

/*
从一个数组中取出最大的数，最小的 求和 和平均值
*/

func main()  {
	var number[5]int=[5]int{1,2,3,4,5}
	var max int=number[0]
	var min int=number[0]
	var sum int
	for i:=0;i<len(number);i++{
		if number[i]>max{
			max = number[i]
		}
		if number[i]<min{
			min= number[i]
		}
		sum = sum+number[i]
	}
	fmt.Println("max",max)
	fmt.Println("min",min)
	fmt.Println("sum",sum)
	fmt.Printf("av: %.2f",float64(sum)/5)
}
