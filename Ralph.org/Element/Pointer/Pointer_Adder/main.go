/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [python//GO//json//Yaml//Other]
 * @Date: 2020-09-09 07:07:01
 * @LastEditTime: 2020-09-09 08:00:09
 * @FilePath: /go/src/github.com/Ralph.org/Element/Pointer/Pointer_Adder/main.go
 * @Effect: DO
 */
package main

import "fmt"

/*Notes：
指针：
指正就是地址，指针变量就是存储地址的变量
*p : 解引用、间接引用

栈帧:   用来给函数运行提供内存空间。 取内存于 stack 上。
		当函数调用时，产生栈帧。函数调用结束，释放栈帧
		栈帧存储： 1. 局部变量。 2. 形参。 （形参与局部变量存储地位等同） 3. 内存字段描述值

指针使用注意：
		空指针：未被初始化的指针。	var p *int		*p --> err
		野指针：被一片无效的地址空间初始化。

变量存储：
		等号 左边的变量，代表 变量所指向的内存空间。	    （写）
		等号 右边的变量，代表 变量内存空间存储的数据值。	（读）

指针的函数传参（传引用）。

		传地址（引用）：将形参的地址值作为函数参数传递。
		传值（数据据）：将实参的 值 拷贝一份给形参。	
		传引用：	在A栈帧内部，修改B栈帧中的变量值。	
		函数传参永远只会值传递 实参将自己的值拷贝一份给形参	
*/

func Pointer_fundation() {
	var (
		a int = 10  //Create RAM space
		p *int = &a
	)
	a = 100
	fmt.Println(a)
	// 借助a的变量地址 对a空间惊醒操作
	*p = 250
	fmt.Println(a,*p)

	a = 1000
	fmt.Println(a)
}

func Nil_Pointer() {
	var p *int 
	fmt.Println(*p)
}

func New_Pointer() {
	var p *int
	p = new(int) // Output default value
	*p = 100
	fmt.Println(*p)
	fmt.Printf("%v\n",*p) //output go-string
}

// Pointer as function parameter and return value
func Fun_Pointer() {
	a,b := 10,20
	Swap(&a,&b)
	fmt.Println(a,b)

}
func Swap(a,b *int) {
	*a,*b = *b,*a
}

func main() {
	// Pointer_fundation()
	// Nil_Pointer()
	// New_Pointer()
	Fun_Pointer()
}