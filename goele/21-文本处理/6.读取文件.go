/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2020-04-27 21:59:40
 * @LastEditTime: 2020-08-06 21:54:06
 * @FilePath: \go\goele\21-文本处理\6.读取文件.go
 * @Effect: 读取文件数据
 */

package main

import (
	"fmt"
	"os"
)

/*
 1.打开文件
 2.读取文件
 3.关闭文件

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
	// 读取文件
	buffer := make([]byte, 1024*2) // 存储从文件读出来的数据 并设置大小
	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Println(string(buffer[:n])) //将文件内容转换成AIISC码值 不足的内存用0来表示   :n表示排除0  string 把AIISC码转换成字母
	// 关闭文件
}
