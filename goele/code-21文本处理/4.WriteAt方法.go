package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	file, err := os.Create("c:/Test/d.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.WriteString("Hello World")
	var str string = "aaa"
	num, _ := file.Seek(10, io.SeekEnd) // 将光标定位到文件中原有内容的后面,返回文件中原有数据的长度
	fmt.Println("num=", num)
	n, err := file.WriteAt([]byte(str), num) // 第二个参数：表示在指定位置写入数据。
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

}
