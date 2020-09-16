/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [GO]
 * @Date: 2020-09-16 09:04:58
 * @LastEditTime: 2020-09-16 09:13:44
 * @FilePath: /go/src/github.com/Ralph.org/Element/Struct/Struct_Addr/struct_pointer.go
 * @Effect: DO
 */
package main 

import (
	"unsafe"
	"fmt"
)


type Person struct{
	name string
	sex	byte
	age int
}


func Pointer_struct_example1() {
	var p1 *Person = &Person{"fuck",'m',29}
	var p2 Person = Person{"sb",'m',92}
	var p3 *Person
	p3 = &p2	
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	p4:= new(Person)
	p4.age = 30
	p4.name = "mdzz"
	p4.sex = 'f'

	fmt.Printf("p3, type= %T\n", p4)
	fmt.Println(unsafe.Sizeof(p4))
	fmt.Println(p4)

	test(p4)
	fmt.Println(p4)
}

func test(p *Person) {
	fmt.Println(unsafe.Sizeof(p))
	p.name = "UFO"
	p.age = 2000
	p.sex = 'n'
}


func main() {
	Pointer_struct_example1()
}