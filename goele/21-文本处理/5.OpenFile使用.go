/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2020-04-27 21:59:40
 * @LastEditTime: 2020-08-06 21:45:07
 * @FilePath: \go\goele\21-文本处理\5.OpenFile使用.go
 * @Effect: 向已有的文件写入数据
 */

package main

import (
	"fmt"
	"os"
)

/*
向已有的文件写入数据
OpenFile()
第一个参数：打开文件的路径
第二个参数：O_RDONLY(只读模式)，O_WRONLY(只写模式), O_RDWR( 可读可写模式)，O_APPEND(追加模式)。
第三个参数表示:  权限，取值范围（0-7）
表示如下：
0：没有任何权限
1：执行权限(如果是可执行文件，是可以运行的)
2: 写权限
3: 写权限与执行权限
4： 读权限
5:  读权限与执行权限
6:  读权限与写权限
7:  读权限，写权限，执行权限

*/
// =======================================Main Function start==========================
func main() {
	file, err := os.OpenFile("F:/go/goele/21-文本处理/Test.txt", os.O_APPEND, 6)
	if err != nil {
		fmt.Println(err)
	}
	// 对文件进行操作
	defer file.Close()
	n, err := file.WriteString("\nbaobao")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
