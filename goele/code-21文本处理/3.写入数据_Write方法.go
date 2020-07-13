package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.Create("c:/Test/c.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	var str string = "Hello World"
	n, err := file.Write([]byte(str)) // 需要将字符串转成字节切片
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	//file.WriteString()
}
