package main

import (
	"fmt"

)

type Student struct {
	name    string
	sex     string
	age     int
	chinese float64
	math    float64
	english float64
}

//Define method for struct
func (s *Student) SayHello(userName string, userAge int, userSex string) {
	// Initialization
	s.name = userName
	s.age = userAge
	s.sex = userSex
	// Judgment conditions
	if s.sex != "男" && s.sex != "女" {
		s.sex = "男"
	}
	if s.age < 1 || s.age > 100 {
		s.age = 18
	}
	fmt.Printf("我叫%s，今年%d岁了，性别是%s\n", s.name, s.age, s.sex)
}

func (s *Student) GetScore(chinese float64, math float64, english float64) {
	// 1: Initialization
	s.chinese = chinese
	s.math = math
	s.english = english
	// 2: 进行计算
	sum := s.chinese + s.math + s.english
	// 3： output
	fmt.Printf("我叫%s,总分%f,平均分%f\n", s.name, sum, sum/3)

}

// =========================main function============================
func main() {
	var stu Student
	stu.SayHello("yaya", 1000, "s")
	stu.GetScore(99, 100, 98)
}
