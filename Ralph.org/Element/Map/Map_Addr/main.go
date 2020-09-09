/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [GO]
 * @Date: 2020-09-09 09:46:44
 * @LastEditTime: 2020-09-09 10:12:26
 * @FilePath: /src/github.com/Ralph.org/Element/Map/Map_Addr/main.go
 * @Effect: DO
 */
/*Notes
字典、映射  key —— value	key： 唯一、无序。 不能是引用类型数据。
	************map 不能使用 cap（）
创建方式：
	1.  var m1 map[int]string		--- 不能存储数据
	2. m2 := map[int]string{}		---能存储数据
	3. m3 := make(map[int]string)		---默认len = 0
	4. m4 := make(map[int]string, 10)
初始化：
	1. var m map[int]string = map[int]string{ 1: "aaa", 2:"bbb"}	保证key彼此不重复。
	2. m := map[int]string{ 1: "aaa", 2:"bbb"}
赋值:
	赋值过程中，如果新map元素的key与原map元素key 相同 	——> 覆盖（替换）
	赋值过程中，如果新map元素的key与原map元素key 不同	——> 添加
map的使用：
	遍历map：
		for  key值， value值 := range map {
		} 
		for  key值 := range map {
		}	
	判断map中key是否存在。
		map[下标] 运算：返回两个值， 
			第一个表 value 的值，如果value不存在。 nil
			第二个表 key是否存在的bool类型。存在 true， 不存在false
删除map：
	delete()函数： 	参1： 待删除元素的map	参2： key值
	delete（map， key）	删除一个不存在的key ， 不会报错。
	map 做函数参数和返回值，传引用。
*/

package main

import "fmt"

func Judg_map() {
	var m9 map[int]string = map[int]string{1:"Luffy", 130:"Sanji", 1301:"Zoro"}

	if v, has := m9[12]; has {	// m9[下标] 返回两个值，第一个是value，第二个是bool 代表key是否存在。
		fmt.Println("value=", v, "has=", has)
	} else {
		fmt.Println("false value=", v, "has=", has)
	}
}


func mapDelete(m map[int]string, key int) {
	delete(m,key)
}


func main() {
	// Judg_map()
	m := map[int]string{1:"Luffy", 130:"Sanji", 1301:"Zoro"}
	fmt.Println("before delete m :", m)

	mapDelete(m, 130)

	fmt.Println("after delete m :", m)

}

/*
package main

import (
	"strings"
	"fmt"
)

func wordCountFunc(str string) map[string]int {
	s := strings.Fields(str)		// 将字符串，拆分成 字符串切片s
	m := make(map[string]int)		// 创建一个用于存储 word 出现次数的 map

	// 遍历拆分后的字符串切片
	for i:=0; i<len(s);i++ {
		if _, ok := m[s[i]]; ok {			// ok == ture 说明 s[i] 这个key存在
			m[s[i]] = m[s[i]]+1			// m[s[i]]++
		} else {							// 说明 s[i] 这个key不存在， 第一次出现。添加到map中
			m[s[i]] = 1
		}
	}
	return m
}

func wordCountFunc2(str string) (m map[string]int) {
	m = make(map[string]int)
	arr := strings.Fields(str)
	for _, v := range arr {
		m[v]++
	}
	return
}

func main()  {
	str := "I love my work and I I I I love love love my family too"
	//mRet := wordCountFunc(str)

	mRet := wordCountFunc2(str)

	// 遍历map ，展示每个word 出现的次数：
	for k, v := range mRet {
		fmt.Printf("%q:%d\n", k, v)
	}
}


*/