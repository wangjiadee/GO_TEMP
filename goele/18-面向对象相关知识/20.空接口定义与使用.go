/*======================================================================
author：Ralph
time：2020-07-21
effect：空接口定义
======================================================================*/
package main

import "fmt"

// =========================main function============================
func main() {
	var i interface{}
	i = 520
	i = "yaya"
	fmt.Println(i) //输出是最后的yaya
	var s []interface{}
	s = append(s, 123, "yaya", 0.333)
	for j := 0; j < len(s); j++ {
		fmt.Println(s[j])
	}
}
