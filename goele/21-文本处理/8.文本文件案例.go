/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2020-04-27 21:59:40
 * @LastEditTime: 2020-08-06 22:09:36
 * @FilePath: \go\goele\21-文本处理\8.文本文件案例.go
 * @Effect: 文件操作的案例--文件拷贝
 */
package main

import (
	"fmt"
	"io"
	"os"
)

// =======================================Main Function start==========================
func main() {
	// 1：打开原有的文件
	file, err := os.Open("F:/go/goele/21-文本处理/Test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// 2：创建新的文件
	file2, err := os.Create("F:/go/goele/21-文本处理/Test2.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	defer file2.Close()

	// 3：将原有的文件读取出来写入到新的文件
	buffer := make([]byte, 1024*2)
	for {
		n, err := file.Read(buffer)
		// 光标在末尾
		if err == io.EOF {
			break
		}
		file2.Write((buffer[:n]))
	}
	// 4：关闭文件
}
