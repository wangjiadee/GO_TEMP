package main

import "fmt"

func main() {
	var m map[int]string=map[int]string{1:"张三",2:"李四",3:"王五",4:"itcast"}  // key是唯一的。
	m1 := map[int]string{1:"张三",2:"李四",3:"王五",4:"itcast"}
	m2 := make(map[string]int,10)
	m2["张三"]=12
	m2["李四"]=15
	m2["张三"]=16  //完成数据的修改。
	fmt.Println(m)
	fmt.Println(m1)
	fmt.Println(len(m2)) // len()返回的是map中已有的键值对个数。
	fmt.Println(m2)
}
