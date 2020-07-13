package main

import "fmt"

//func main() {
//	var qiepian[]int
//	fmt.Println(qiepian)
//	fmt.Println(len(qiepian))
//}
//自动推导方式定义切片
//func main(){
//	s := []int{}
//	fmt.Println(s)
//	fmt.Println(len(s))
//}

//func main()  {
//	s:=make([]int,3,5)
//	fmt.Println(s)
//	fmt.Println(len(s))
//	fmt.Println(cap(s))
//}

//切片的初始化
//func main() {
//	var s []int
//	s = append(s, 1, 2, 3, 4, 5, 6)
//	fmt.Println(s)
//	s[3]=33
//	fmt.Println(s)
//}

//func main() {
//	s:=[]int{77,88,99,100}
//	s = append(s,101,102)
//	fmt.Println(s)
//}

//func main() {
//	s:= make([]int,3,10)
//	s[1]=20
//	fmt.Println(s)
//}

//循环初始化：
//func main()  {
//	s:=make([]int,3,10)
//	for i:=0;i<len(s) ;i++  {
//		s[i]=i
//	}
//	s = append(s,80)
//	s[3]=70
//	fmt.Println(s)
//}

//切片的循环遍历
//func main() {
//	s:= []int{1,2,3,4,5,6,7,8}
//	for i:=0;i<len(s) ;i++  {
//		fmt.Println(s[i])
//	}
//}

//func main() {
//	//s:= []int{1,2,3,4,5,6}
//	s:= make([]int,4,5)
//	for i,v:=range s{
//		fmt.Println("i=",i)
//		fmt.Println("v=",v)
//	}
//}

//func main() {
//	s:= []int{1,2,3,4,5,6}
//	// 0 起始位 3 终止位 5 容量  也是顾头不顾尾
//	//容量=第三个-第一个
//	//长度=第二个-第一个
//	s1:= s[0:3:5]
//	fmt.Println(s1)
//	fmt.Println(cap(s1))
//	fmt.Println(len(s1))
//	s1 := s[:]
//	fmt.Println(s1)
//	fmt.Println(len(s1))
//	fmt.Println(cap(s1))
//	s2:=s[2:]
//	fmt.Println(s2)
//	fmt.Println(len(s2))
//	fmt.Println(cap(s2))
//	s3:= s[:3]
//	fmt.Println(s3)
//	fmt.Println(len(s3))
//	fmt.Println(cap(s3))
//}

//func main() {
//	s:= []int{1,2,2,3,4,5,6,7,8,9}
//	//s[0]=30
//	s1:=s[2:5]
//	s1[0] = 88
//	fmt.Println(s1)
//	fmt.Println(s)
//}

//关于append基本使用
//func main(){
//	s:=make([]int,5,8)
//	s = append(s,1)
//	s = append(s,2)
//	s = append(s,3)
//	fmt.Println(s)
//	fmt.Println(len(s))
//	fmt.Println(cap(s))
//
//}

//func main()  {
//	s1:=[]int{1,2}
//	s2:=[]int{4,3,2,1}
//	//copy(s1,s2)
//	copy(s2,s1)
//	fmt.Println(s2)
//}

//函数
//	func main() {
//		//s:=[]int{1,2,3,4,5,6}
//		s:= make([]int,10)
//		Init(s)
//		fmt.Println(s)
//	}
//	func Init(num[]int)  {
//		for i:=0;i<len(num) ;i++  {
//			num[i] =i
//		}
//	}

	func main() {
		var count int
		fmt.Println("输入求和的个数")
		fmt.Scan(&count)
		s := make([]int, count)
		InitD(s)
		SumAdd(s)
		sum := SumAdd(s)
		fmt.Println("sum", sum)
	}
	func InitD(num []int) {
		for i := 0; i < len(num); i++ {
			fmt.Println("请输入第%d个整数\n", i+1)
			fmt.Scan(&num[i])
		}
	}
	func SumAdd(num []int) int {
		var sum int
		for i := 0; i < len(num); i++ {
			sum += num[i]
		}
		return sum
	}

	func main()  {
		var count int
		//1明确格数
		fmt.Println("请输入比较的数据个数：")
		fmt.Scan(&count)
		//2 输入数据
		s:= make([]int,count)
		AddData(s)
		max:= compareValue(s)
		fmt.Println("max=",max)
	}
	func AddData(num []int)  {
		for i:=0;i<len(num) ;i++  {
			fmt.Printf("请输入第%d个数字",i+1)
			fmt.Scan(&num[i])
		}
	}

	func compareValue(num[]int) int {
		var max int = num[0]
		for i:=0;i<len(num) ;i++  {
			if num[i] > max{
				max = num[i]
			}
		}
		return max
	}