package main

import "fmt"

type Student struct {
	Person
	score float64
}
type Person struct {
	id int
	name string
	age int
}

func(p *Person) PrintInfo()  {
	fmt.Println(*p)
}

func main() {
  stu:=Student{Person{101,"张三",18},90}
  stu.PrintInfo()
}
