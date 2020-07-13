package main

import "fmt"

func main() {
	var p[2]*int
	var i int = 10
	var j int = 20
	p[0]=&i
	p[1]=&j
	fmt.Println(p) // 打印的结构都是地址
	fmt.Println(*p[0]) // 不要加括号
	fmt.Println(*p[1])
	//循环
	//for i:=0;i<len(p);i++{
	//	fmt.Println(*p[i])
	//}
	for k,value:= range p{
		fmt.Println(k)
		fmt.Println(*value)
	}

}
