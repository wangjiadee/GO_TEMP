package main

import "fmt"

func TestSum(args...int){
	// args参数可以认为就是一个集合，在该集合中存储了传递过来的数据，
	// 会对每个数据加上一个编号，该编号是从0开始计算的。所以可以通过该编号获取集合中存储的具体的数据。
	/*
	fmt.Println(args[0])
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])
	*/
	/*
	for i:=0;i<len(args);i++{ // len() 可以获取参数的存储的数据的个数
		fmt.Println(args[i])
	}
	*/
	// 集合。 数组，切片
	// 如果只写一个变量，该变量中存储的是集合的编号
	// _:匿名变量。匿名变量不会保存具体的数据
	for _,v:= range args{
		//fmt.Println("a=",a) // 存储的是集合的编号
		fmt.Println("v=",v) // 存储的是具体的值。
	}

}
// 固定参数放在前面，不定参数放在后面
func TestA(num int,args...int){
	fmt.Println(num)
	fmt.Println("argsp[0]=",args[0])
	fmt.Println("argsp[1]=",args[1])
}
func main() {
  //TestSum(3,6,7,10)
  TestA(10,8,9)
}

