/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2020-04-27 21:59:40
 * @LastEditTime: 2020-08-06 20:10:02
 * @FilePath: \go\goele\21-文本处理\1.创建文件.go
 * @Effect: GO文本文件处理
 */

package main

import (
	"fmt"
	"os"
)

/*
 导入os包 创建文件
 执行create（） 镜像文件的创建 最后要关闭文件
*/
// =======================================Main Function start==========================
func main() {
	// 注意斜线的符号，同时指定文件的路径和文件名称
	file, err := os.Create("F:/go/goele/21-文本处理/Test.txt")
	if err != nil {
		fmt.Println(err)
		// file.Close()
	}
	// 延迟执行关闭文件句柄
	defer file.Close()
	// 关闭文件
	file.Close()
}
