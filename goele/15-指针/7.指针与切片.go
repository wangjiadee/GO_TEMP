package main

import "fmt"

func main() {
	s:=[]int{1,2,3,4,5,6}
	var p *[]int
	p=&s
	//fmt.Println(*p)
	fmt.Println((*p)[0])
	//fmt.Println(p[0]) //错误
	//(*p)[0]=200
	//fmt.Println(s)


	//for i:=0;i<len(*p) ; i++ {
     // fmt.Println((*p)[i])
	//}

	for k,value:= range *p{
		fmt.Println("k=",k)
		fmt.Println("value=",value)
	}
}
