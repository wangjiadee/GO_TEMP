/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [GO]
 * @Date: 2020-09-16 09:21:23
 * @LastEditTime: 2020-09-16 09:28:10
 * @FilePath: /go/src/github.com/Ralph.org/Element/Struct/Struct_Addr/struct_para.go
 * @Effect: DO
 */
package main

import "fmt"

type Person struct {
	name string
	age int
	flg bool 
	intereset []string
}

//  Initializing the structure by function parameters 
func initFunc(p *Person) {
	p.name = "naine"
	p.age = 35
	p.flg = true
	p.intereset = append(p.intereset,"fucking u")
	p.intereset = append(p.intereset,"fucking m")
}

// Initializes the structure through the function return value
func initFunc2() *Person {
	p := new(Person)
	p.name = "wuliy"
	p.age = 40
	p.flg = false
	p.intereset = append(p.intereset,"fucking e")
	p.intereset = append(p.intereset,"fucking b")
	return p   //返回指针变量的值 -heap的地址
}


func main() {
	var p1 Person
	initFunc(&p1)
	fmt.Println(p1)

	p2 := initFunc2()
	fmt.Println(p2)
}