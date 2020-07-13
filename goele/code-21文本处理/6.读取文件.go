package main

import (
	"os"
	"fmt"
)

func main() {
	// 1: 打开要读取的文件
	file, err := os.Open("c:/Test/file.txt") // 注意：OpenFile方法的区别
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// 2: 进行文件内容读取
	// 定义一个字符类型切片，存储从文件中读取的数据
	buffer := make([]byte, 1024*2) // 大小为2kb
	n, err := file.Read(buffer) // 将从文件中读取的数据保存到字符切片中， n:表示从文件中读取的数据的长度
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buffer[:n]))
	fmt.Println("读取数据的长度:",n)
	// 3: 关闭文件
}
