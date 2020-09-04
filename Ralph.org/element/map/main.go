/*
 * @Author: Ralph
 * @Date: 2020-09-02 09:46:14
 * @LastEditTime: 2020-09-04 11:47:33
 * @LastEditors: Please set LastEditors
 * @Description: GO的map介绍和案例
 * @FilePath: go\src\github.com\Ralph.org\element\go_package_method\cala\tempCodeRunnerFile.go
 */
package main

import "fmt"

// Create map and initialization
// Like python dict
func C_map() {
	// three initialization method
	var m map[int]string = map[int]string{1: "zs", 2: "ls", 3: "ww", 4: "ss"}
	m1 := map[int]string{1: "aaa", 2: "bbb", 3: "cccc"}
	m2 := make(map[string]int, 10)
	fmt.Println(m, m1, m2)
	// change map value
	m2["zs"] = 12
	m2["1"] = 22
	fmt.Println(m2)
}

// About map key or value
func map_kv_l() {
	var m map[int]string = map[int]string{1: "sb", 2: "mdzz"}
	fmt.Println(m[2])

	v, success := m[5]
	if success {
		fmt.Println(v)
	} else {
		fmt.Println(v)
		fmt.Println("Error: value not exits")
	}
}

func Traverse_map() {
	var m map[int]string = map[int]string{1: "sb", 2: "mdzz"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func Delete_map() {
	var m map[int]string = map[int]string{1: "zs", 2: "ls", 3: "ww", 4: "ss"}
	delete(m, 2)
	fmt.Println(m)
}

func Map_Pass_parameters() {
	var m map[string]string = map[string]string{"json": "jksla", "python": "kzkz"}
	DeleteMap(m)
	PrintMap(m)
}

func PrintMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
func DeleteMap(m map[string]string) {
	delete(m, "json")
	fmt.Println(m)
}

func MapExample() {
	str := "mdzzoppo"
	m := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		ch := str[i]
		// fmt.Println(ch)
		m[ch] = m[ch] + 1
	}
	for k, v := range m {
		fmt.Printf("%c:%d\n", k, v)
	}
}

// #####################start main function
func main() {
	C_map()
	map_kv_l()
	Traverse_map()
	Delete_map()
	Map_Pass_parameters()
	MapExample()
}

