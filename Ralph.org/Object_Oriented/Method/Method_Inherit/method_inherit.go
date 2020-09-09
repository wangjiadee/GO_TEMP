package main

import (
	"fmt"

)

//Sub Class
type Rep struct {
	Person
	Hobby string // 爱好
}

// Sub class Method
func (r *Rep) RepSayHello(h string) {
	r.Hobby = h
	fmt.Printf("我叫%s ，我的爱好是%s，我的年龄是%d，我是一个%s狗仔\n", r.name, r.Hobby, r.age, r.sex)
}

// Sub Class
type Pro struct {
	Person
	WorkYear int
}

// Sub class
func (p *Pro) ProSayHello(workYear int) {
	p.WorkYear = workYear
	fmt.Printf("我叫%s，我的年龄是%d，我是%s，我的工作年限是 %d年\n", p.name, p.age, p.sex, p.WorkYear)
}

// futher Class
type Person struct {
	name string
	age  int
	sex  string
}

// futher Class method
func (p *Person) SetValue(userName string, userAge int, userSex string) {
	p.sex = userSex
	p.age = userAge
	p.name = userName
}

// =========================main function============================
func main() {
	var rep Rep
	var pro Pro
	rep.SetValue("YAYA", 14, "w")
	rep.RepSayHello("吃饭")
	pro.SetValue("yy", 98, "girl")
	pro.ProSayHello(5)
}

