/*======================================================================
author：Ralph
time：2020-07-21
effect：go方法继承
======================================================================*/
package main

import (
	"fmt"
)

// 子类
type Rep struct {
	Person
	Hobby string // 爱好
}

// 子类的方法
func (r *Rep) RepSayHello(h string) {
	r.Hobby = h
	fmt.Printf("我叫%s ，我的爱好是%s，我的年龄是%d，我是一个%s狗仔\n", r.name, r.Hobby, r.age, r.sex)
}

type Pro struct {
	Person
	WorkYear int
}

func (p *Pro) ProSayHello(workYear int) {
	p.WorkYear = workYear
	fmt.Printf("我叫%s，我的年龄是%d，我是%s，我的工作年限是 %d年\n", p.name, p.age, p.sex, p.WorkYear)
}

// 父类
type Person struct {
	name string
	age  int
	sex  string
}

// 父类的方法
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
	pro.SetValue("yy", 98, "nv")
	pro.ProSayHello(5)
}
