/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2019-03-04 16:11:06
 * @LastEditTime: 2020-08-10 22:08:13
 * @FilePath: \go\goele\22-字符串处理\1.字符串处理.go
 * @Effect: go 的字符串方法
 */

package main

import (
	"fmt"
	"strings"
)

/*
Notes：

Contains :  判断字符串s是否包含了某个子字符
			存在返回T，不存在返回F
			func Contains(s, substr string) bool

join：		将一系列字符串连接为一个字符串，之间用sep来分隔。

index:		在某个字符串中查找某个字符串的位置
			步长从0开始

Repeat ：   某个字符串重复多少次，返回的是重复后的字符串

Replace：   在s字符串中，把old字符串替换为new字符串，
			n表示替换的次数，小于0表示全部替换

Split  ：   把s字符串按照sep分割，返回slice（切片）


@@LINK@@:	https://studygolang.com/pkgdoc
*/

// =======================================Main Function start==========================
func main() {
	var str string = "xiejier"
	b := strings.Contains(str, "xie")
	fmt.Println(b)

	s := []string{"Oppo", "vivo", "huawei"}
	str := strings.Join(s, "--")
	fmt.Println(str)

	str1 := "wine"
	n := strings.Index(str1, "n")
	fmt.Println(n)

	s := strings.Repeat("yaya", 5)
	fmt.Println(s)

	str := "oppo find x"
	str3 := strings.Replace(str, "x", "x2", 1)
	fmt.Println(str3)

	var str4 string = "mtk@mo@vivo"
	str5 := strings.Split(str4, "@")
	fmt.Println(str5)

}
