package main

//记者：我叫张三 ，我的爱好是偷拍，我的年龄是34，我是一个男狗仔。
//程序员：我叫孙全，我的年龄是23，我是男生，我的工作年限是 3年。
// 1: 定义父类
/*
type Person struct {
	name string
	age  int
	sex  string
}

// 2: 给父类添加方法。
func (p *Person) SetValue(userName string, userAge int, userSex string) {
	p.sex = userSex
	p.age = userAge
	p.name = userName
}

// 3: 定义相应的子类。
// 记者类
type Rep struct {
	Person
	Hobby string // 爱好
}

func (r *Rep) RepSayHello(h string) {
	r.Hobby = h
	fmt.Printf("我叫%s ，我的爱好是%s，我的年龄是%d，我是一个%s狗仔\n", r.name, r.Hobby, r.age, r.sex)
}

// 程序员类
type Pro struct {
	Person
	WorkYear int
}

func (p *Pro) ProSayHello(workYear int) {
	p.WorkYear = workYear
	fmt.Printf("我叫%s，我的年龄是%d，我是%s，我的工作年限是 %d年\n",p.name,p.age,p.sex,p.WorkYear)
}

// 4: 给子类添加相应的方法
// 5: 展示效果
*/
func main() {
	//var rep Rep
	//rep.SetValue("张三", 26, "男")
	//rep.RepSayHello("偷拍")
	//
	//var pro Pro
	//pro.SetValue("李四", 28, "男")
	//pro.ProSayHello(3)

}
