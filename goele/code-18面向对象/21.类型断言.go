package main

import "fmt"

func main() {
	var i interface{}
	i="abc"
	value,ok:=i.(int)
	if ok{
		fmt.Println(value+10)
	}else{
		fmt.Println("类型推断错误")
	}

}
