/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2020-04-27 21:59:40
 * @LastEditTime: 2020-08-06 22:01:53
 * @FilePath: \go\goele\21-文本处理\7.循环读取文件数据.go
 * @Effect: 循环读取数据
 */
package main

import (
	"fmt"
	"io"
	"os"
)

/*


 */
// =======================================Main Function start==========================
func main() {
	// 打开文件
	// Open 调用Openfile 只读无权限的方式打开
	file, err := os.Open("F:/go/goele/21-文本处理/Test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buffer := make([]byte, 10)
	for {
		n, err := file.Read(buffer)
		// 光标在末尾
		if err == io.EOF {
			break
		}
		fmt.Print(string(buffer[:n]))
	}

}
