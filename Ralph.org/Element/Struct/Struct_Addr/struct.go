/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [GO]
 * @Date: 2020-09-16 08:36:46
 * @LastEditTime: 2020-09-16 09:00:48
 * @FilePath: /go/src/github.com/Ralph.org/Element/Struct/Struct_Addr/struct.go
 * @Effect: 结构体加强
 */


package main
import (
	"unsafe"
	"fmt"
)
// // 类型定义 等价于（int byte bool string）
type Person struct {
	name string
	sex byte
	age int
}

func test(man Person) {
	// 求一个变量的大小
	fmt.Println(unsafe.Sizeof(man))
	man.name = "fucking"
	man.age = 33
	fmt.Println(man)
}





func main() {
	//  Structure definition and initialization
	// 顺序初始化
	var man Person = Person{"Ralph",'m',22}
	fmt.Println(man)
	// 指定初始化
	man1 := Person{name:"rose",age:22}
	fmt.Println(man1)

	// Index member variable 
	man2 := Person{sex:'f',age:22}
	fmt.Printf("man2.name = %q\n",man2.name) 

	var man3 Person
	man3.name = "Jon"
	man3.age = 8
	man3.sex = 'f'
	fmt.Println(man3)
	man3.age = 43
	fmt.Println(man3)

	// Comparison of structures 
	var p1 Person = Person{"tom",'m',9}
	var p2 Person = Person{"tom",'m',19}
	var p3 Person = Person{"tom",'m',29}

	fmt.Println(p1 == p2)
	fmt.Println(p1 == p3)

	// Assignment of structures of the same type 
	var temp Person
	fmt.Println(temp)
	temp = p3
	fmt.Println(temp)

	// The structure body is used to transfer parameters inside the function
	var t Person
	fmt.Println(unsafe.Sizeof(t))
	test(t)
	fmt.Println(t)
	fmt.Println(&t)
	fmt.Println(t.name)
	fmt.Println(unsafe.Sizeof(true))
}




