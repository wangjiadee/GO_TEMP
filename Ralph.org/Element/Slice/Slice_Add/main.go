/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [GO]
 * @Date: 2020-09-09 08:12:36
 * @LastEditTime: 2020-09-09 09:21:44
 * @FilePath: /go/src/github.com/Ralph.org/Element/Slice/Slice_Add/main.go
 * @Effect: 切片强化
 */
package main

/*Notes:
切片：

为什么用切片：
	1. 数组的容量固定，不能自动拓展。
	2. 值传递。 数组作为函数参数时，将整个数组值拷贝一份给形参。
	在Go语言当，我们几乎可以在所有的场景中，使用切片替换数组使用

切片的本质：
	不是一个数组的指针，是一种数据结构体，用来操作数组内部元素。
	和数组的区别是 数组[]内要添加数字，切片为空或者[...]
	截取数组，初始化切片时，没有指定切片容量时， 切片容量跟随原数组（切片）。

切片的使用：
	切片名称[low:high:max]
	high： len = high -low
	容量： cap = max - low

s[:high:max] : 从 0 开始，到high结束。（不包含）
s[low:] :	从low 开始，到末尾
s[: high]:	从 0 开始，到high结束。容量跟随原先容量。(常用)


append：在切片末尾追加元素
	append(切片对象， 待追加元素）
	向切片增加元素时，切片的容量会自动增长。1024 以下时，一两倍方式增长。
copy：
	copy（目标位置切片，源切片）
	拷贝过程中，直接对应位置拷贝。
*/

import (
	"fmt"
)

func Define_Slie() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr[1:5:7]
	fmt.Println(s, len(s), cap(s))
	s2 := s[:5]
	fmt.Println(s2, len(s2), cap(s2))
}

func Create_slice() {
	slice := []int{1, 2, 3}
	slice1 := make([]int, 3, 6)
	slice2 := make([]int, 4)
	fmt.Println(slice, slice1, slice2)
}

// make only create slice、map、channel

func Slice_Append() {
	s1 := []int{1, 2, 3, 4, 5}
	s1 = append(s1, 250)
	s1 = append(s1, 250)
	s1 = append(s1, 250)
	s1 = append(s1, 250)
	fmt.Println(s1)
}

func Slice_Example(data []string) []string {
	out := data[:0]
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}
	return out
}

// remove duplicate string
func Remove_slice(new []string) []string {
	res := new[:0]
	for i := range new {
		flag := true
		for j := range res {
			if new[i] == res[j] {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, new[i])
		}
	}
	return res
}

func Copy_slice() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice1, slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置)
	fmt.Println(slice1, slice2)
}

func main() {
	// Define_Slie()
	// Create_slice()
	// Slice_Append()

	// data := []string{"sds","rrf","","sgfh",""}
	// new := []string{"asd","asd","ert","ert","cvx"}
	// l := Slice_Example(data)
	// fmt.Println(l)
	// l2 := Remove_slice(new)
	// fmt.Print(l2)

	Copy_slice()
}
