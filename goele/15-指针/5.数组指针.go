package main

import "fmt"

func main() {
	nums:=[10]int{1,2,3,4,5,6,7,8,9,10}
	var p *[10]int
	p=&nums
	// fmt.Println(*p) // 获取整个数组中的内容。
	// fmt.Println((*p)[3])// []的运算优先级高于*
	// fmt.Println(p[0])
	//for i:=0;i<len(p) ;i++  {
	//	fmt.Println(p[i])
	//}

	UpdateArr(p)
	fmt.Println(nums)
}
func UpdateArr(p*[10]int)  {
	p[0]=100
}
