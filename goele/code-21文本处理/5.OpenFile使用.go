package main

import (
	"os"
	"fmt"
)

//OpenFile()这个函数有三个参数：
//第一个参数表示：打开文件的路径
//第二个参数表示：模式，常见的模式有
//O_RDONLY(只读模式)，O_WRONLY(只写模式), O_RDWR( 可读可写模式)，O_APPEND(追加模式)。
//
//第三个参数表示:  权限，取值范围（0-7）
//表示如下：
//0：没有任何权限
//1：执行权限(如果是可执行文件，是可以运行的)
//2: 写权限
//3: 写权限与执行权限
//4： 读权限
//5:  读权限与执行权限
//6:  读权限与写权限
//7:  读权限，写权限，执行权限
func main() {
	file, err := os.OpenFile("c:/Test/c.txt", os.O_APPEND, 6)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// 通过文件指针向文件中写入数据，读取数据
	n, err := file.WriteString("传智播客")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

}
