package main

import (
	"fmt"
)

//定义一个学生类,有六个属性,分别为姓名、性别、年龄、语文、数学、英语成绩
//第一方法：打招呼的方法：介绍自己叫XX，今年几岁了。是男同学还是女同学。
//第二个方法：计算总分与平均分的方法

// 定义结构体1
type Student struct {
	name    string
	sex     string
	age     int
	chinese float64
	math    float64
	english float64
}

// 为定义好的结构体添加方法
func (s *Student) SayHello(userName string, userAge int, userSex string) {
	// 初始化
	s.name = userName
	s.age = userAge
	s.sex = userSex
	// 判断初始化的值进行判断
	if s.sex != "男" && s.sex != "女" {
		s.sex = "男"
	}
	if s.age < 1 || s.age > 100 {
		s.age = 18
	}
	fmt.Printf("我叫%s，今年%d岁了，性别是%s\n", s.name, s.age, s.sex)
}

func (s *Student) GetScore(chinese float64, math float64, english float64) {
	// 1: 初始化
	s.chinese = chinese
	s.math = math
	s.english = english
	// 2: 进行计算
	sum := s.chinese + s.math + s.english
	// 3： 打印输出结果
	fmt.Printf("我叫%s,总分%f,平均分%f", s.name, sum, sum/3)

}

// =========================main function============================
func main() {
	var stu Student
	stu.SayHello("yaya", 1000, "s")
	stu.GetScore(99, 100, 98)
}
