// recover 只有在defer调用中的函数生肖  defer 表示延迟函数
package main

import "fmt"

func Test(n int) {
	defer TestRecover()
	var num [10]int
	num[n] = 12
	fmt.Println(num)
	fmt.Println("aaaaa")

}
func TestRecover() {
	recover()              // 如果报错 将不显示东西
	fmt.Println(recover()) //如果想看到输出的内容 runtime error: index out of range [11] with length 10
}

// =======================================Main Function start==========================
func main() {
	Test(11)
}
