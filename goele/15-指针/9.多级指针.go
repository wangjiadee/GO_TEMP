package main

import "fmt"

func main() {
var a int = 10
var p *int //一级指针
p =&a

var pp **int
pp =&p
**pp=200
fmt.Println(a)

}
