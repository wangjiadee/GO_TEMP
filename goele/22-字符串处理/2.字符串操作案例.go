/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2019-03-04 16:11:06
 * @LastEditTime: 2020-08-10 22:08:52
 * @FilePath: \go\goele\22-字符串处理\2.字符串操作案例.go
 * @Effect: GO 字符串操作案例
 */

package main

import (
	"fmt"
	"strings"
)

/*
Notes：


*/

// =======================================Main Function start==========================
func main() {
	// EG1: 让用户输入一个日期格式，如:2018-01-02, 输出日期为2008年1月2日
	var User_Input string
	fmt.Println(">>>>Pls input time (2020-02-02):")
	fmt.Scan(&User_Input)

	s := strings.Split(User_Input, "-")
	fmt.Println(s[0] + "年" + s[1] + "月" + s[2] + "日")

	// EG2: 让用户输入一句话,判断这句话中有没有“傻逼”,如果有“傻逼”就替换成“**”，然后输出.
	var User_Input string
	fmt.Println(">>>>Pls input something:")
	fmt.Scan(&User_Input)

	if strings.Contains(User_Input, "sb") {
		User_Input = strings.Replace(User_Input, "sb", "**", -1)
	}
	fmt.Println(User_Input)

	// EG3: 从Email中提取出用户命和域名：abc@email.com
	var str4 string = "Ralph@huawei.com"
	str5 := strings.Split(str4, "@")
	fmt.Println(str5)

	// EG4: User输入一句话，找出所有e的位置
	var User_Input string
	fmt.Println(">>>Pls input somethins:")
	fmt.Scan(&User_Input)
	n := strings.Index(User_Input, "e")
	fmt.Println(n)
}
