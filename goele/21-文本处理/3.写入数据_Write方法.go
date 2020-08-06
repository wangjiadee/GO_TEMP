/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2020-04-27 21:59:40
 * @LastEditTime: 2020-08-06 20:24:01
 * @FilePath: \go\goele\21-文本处理\3.写入数据_Write方法.go
 * @Effect: Write 写入数据
 */

package main

import (
	"fmt"
	"os"
)

/*
Write 写入数据
WriteString 相当于把Write 自动的转化成了切片
*/
// =======================================Main Function start==========================
func main() {
	// 注意斜线的符号，同时指定文件的路径和文件名称
	file, err := os.Create("F:/go/goele/21-文本处理/Test.txt")
	if err != nil {
		// 输出文件错误信息
		fmt.Println(err)
	}
	// 延迟执行关闭文件句柄
	defer file.Close()
	var str string = "Fucking u everyday"
	n, err := file.Write([]byte(str)) //need str ---切片
	if err != nil {
		// 输出文件错误信息
		fmt.Println(err)
	}
	fmt.Println(n)
}
