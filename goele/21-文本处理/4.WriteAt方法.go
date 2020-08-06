/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2020-04-27 21:59:40
 * @LastEditTime: 2020-08-06 20:34:46
 * @FilePath: \go\goele\21-文本处理\4.WriteAt方法.go
 * @Effect: WriteAt 写入方法
 */

package main

import (
	"fmt"
	"io"
	"os"
)

// =======================================Main Function start==========================
func main() {
	file, err := os.Create("F:/go/goele/21-文本处理/Test.txt") // 注意斜线的符号，同时指定文件的路径和文件名称
	if err != nil {
		fmt.Println(err) // 输出文件错误信息
	}

	defer file.Close() // 延迟执行关闭文件句柄
	file.WriteString("Fucking u")

	var str string = "YAYA"
	num, _ := file.Seek(0, io.SeekEnd) // 将光标定位到文件中原有内容的后面,返回文件中原有数据的长度 第一个参数表示间距
	fmt.Println("Byte:", num)

	n, err := file.WriteAt([]byte(str), num) // 第二个参数：表示在指定位置写入数据。
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n) //writeAt 指定的位置去写入，同时会覆盖YAYAing u

}
