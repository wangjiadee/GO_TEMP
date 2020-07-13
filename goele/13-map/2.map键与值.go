package main

import "fmt"

func main() {
	// map 类似于python的字典
	var m map[int]string = map[int]string{1: "王五", 2: "李四"}
	// 指定输出类型
	fmt.Println(m[2])

	value, ok := m[6]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println(value)
		fmt.Println("不存在")
	}

	// 循环遍历map
	/*
		for key, value := range m {
			fmt.Println(key)
			fmt.Println(value)
		}
	*/

	// map 的删除操作
	delete(m, 2)
	fmt.Println(m)
}
