/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2020-04-27 21:59:40
 * @LastEditTime: 2020-08-06 20:16:43
 * @FilePath: \go\goele\21-文本处理\2.WriteString方法.go
 * @Effect: WriteString 写入数据
 */

package main

import (
	"fmt"
	"os"
)

/*
WriteString 写入数据
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
	// 写入数据
	N, err := file.WriteString("Fucking u!")
	if err != nil {
		// 输出文件错误信息
		fmt.Println(err)
	}
	fmt.Println(N)
	// 关闭文件
	file.Close()
}
