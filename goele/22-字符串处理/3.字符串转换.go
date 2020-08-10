/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2019-03-04 16:11:06
 * @LastEditTime: 2020-08-10 22:19:03
 * @FilePath: \go\goele\22-字符串处理\3.字符串转换.go
 * @Effect: 字符串转换
 */

package main

import (
	"fmt"
	"strconv"
)

/*
 Notes：


*/

// =======================================Main Function start==========================
func main() {
	// 将bool---> int
	str := strconv.FormatBool(true)
	str1 := strconv.FormatBool(false)
	fmt.Println(str) //输出是字符串
	fmt.Println(str1)

	// 将int--->str
	str2 := strconv.Itoa(123)
	fmt.Println(str2)

	//str---> bool
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b)
	}

	//str ---> int
	num, err := strconv.Atoi("20191109")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}
