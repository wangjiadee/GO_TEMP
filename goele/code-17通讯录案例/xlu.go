package main

import "fmt"

//定义一个联系人和电话的结构体
type Person struct {
	userName string
	// key :表示电话的类型 value 表示电话号码
	addressPhone map[string]string
}

//定义一个多个联系人的切片
var personList = make([]Person, 0)

//添加联系人的函数
func addPerson() {
	var name string
	var address string
	var phone string
	var exit string
	//存储电话的类型和电话
	var addressPhone = make(map[string]string)
	fmt.Println("请输入姓名")
	fmt.Scan(&name)
	//for循环 是为了能存储多个不同类型的电话
	for {
		//存储电话的类型
		fmt.Println("请输入电话类型（座机/手机）")
		fmt.Scan(&address)
		//存储电话的号码
		fmt.Println("请输入电话号码")
		fmt.Scan(&phone)
		//将电话的号码和类型存储到 addressPhone 的map里面
		//key为电话类型，value 为电话号码
		addressPhone[address] = phone
		fmt.Println("如果结束电话的录入，请按Q，继续请按w")
		fmt.Scan(&exit)
		if exit == "Q" {
			break
		} else {
			continue
		}
	}
	//	将整个联系人的信息存储到personlist的切片当中
	personList = append(personList, Person{userName: name, addressPhone: addressPhone})
	//调用函数展示联系人的信息
	showPersonList()
}

//显示联系人的函数
func showPersonList() {
	//判断存储联系人的切片里面有有没有联系人
	if len(personList) == 0 {
		fmt.Println("暂时没有联系人信息")
	} else {
		//	循环遍历出联系人的信息
		for _, v := range personList {
			fmt.Println("联系人姓名：", v.userName)
			for k, v1 := range v.addressPhone {
				fmt.Println("存储电话类型：", k)
				fmt.Println("电话号码：", v1)
			}
		}
	}
}

//删除联系人的信息函数
func deletePerson() {
	//输入删除人的名字--->循环遍历有没有这个人---> 有就删除
	var name string
	// 记录要删除的联系人信息在切片中的下标。 先假定等于-1 因为不知道有没有
	var index int = -1
	fmt.Println("请输入要删除的联系人姓名：")
	fmt.Scan(&name)
	for i := 0; i < len(personList); i++ {
		if personList[i].userName == name {
			index = i
			break
		} else {
			fmt.Println("没有此联系人")
		}
	}
	//如果有删除重新记录到personlist里面
	if index != -1 {
		/* append函数第二个参数如果是切片 后面要给三个点。
		   personlist[:index] = [0:index]
		   personList[inedx+1:] = [index+1:] 到最后
		比如 3 是我们要删除的 那么3的下标就是index 前面是添加了 12 不包括3
		后面是添加了比3的下标多一个的下标 就是4 添加包括4以后的数
		*/
		personList = append(personList[0:index], personList[index+1:]...)
	}
	showPersonList()
}

// 查询联系人
func findPerson() *Person {
	//	根据输入的联系人的名字---> 找对应的信息 ---> 输出打印
	var name string
	var index int = -1
	fmt.Println("请输入要查询的联系人姓名：")
	fmt.Scan(&name)
	//循环personList 大切片的内容 如过有 记录下标 并在存循环
	for k, v := range personList {
		if v.userName == name {
			index = k
			fmt.Println("联系人姓名：", v.userName)
			for key, value := range v.addressPhone {
				fmt.Printf("%s:%s\n", key, value)
			}
		}
	}
	if index == -1 {
		fmt.Println("没有找到联系人信息")
		return nil
	} else {
		return &personList[index]
	}
}

//编辑联系人
func editPerson() {
	//	先找到要编辑的信息(调用findPerson)---编辑表示要影响到原来对应的信息
	//	所以要用到指针
	var name string
	var num int                  // 存储修改的数据的类型
	var menu = make([]string, 0) // 保存要修改的电话类型，方便后面修改
	var pNum int                 // 编辑的电话类型编号
	var phone string             // 新的电话号码
	var p *Person
	p = findPerson()
	//存储新的联系人的姓名
	if p !=nil {
		//死循环：
		for {
			//	修改联系人的姓名
			fmt.Println("编辑用户名称请按:5,编辑电话请按:6,退出请按:7")
			fmt.Scan(&num)
			switch num {
			case 5:
				fmt.Println("请输入新的姓名：")
				fmt.Scan(&name)
				//能影响原来的数据
				p.userName = name
				showPersonList()
			case 6:
				var j int
				//	编辑电话  展示联系人所有的电话信息
				for k,v:=range p.addressPhone {
					fmt.Println("编辑（",k,")",v,"请按：",j)
					menu = append(menu,k)
					j++
				}
				//	完成修改
				fmt.Println("请输入要编辑的号码的类型：")
				fmt.Scan(&pNum)
				for index,v1 := range menu{
					if index == pNum{
						fmt.Println("请输入新的电话号码：")
						fmt.Scan(&phone)
						p.addressPhone[v1]=phone
					}
				}
			}
			if num == 7{
				break
			}
		}


	} else {
		fmt.Println("没有找到要编辑的联系人信息")
	}
}

func main() {
	for {
		scanNum()
	}
}

// 给出相应的操作类型
func scanNum() {
	fmt.Println("添加联系人信息，请按1")
	fmt.Println("删除联系人信息，请按2")
	fmt.Println("查询联系人信息，请按3")
	fmt.Println("编辑联系人信息，请按4")
	//对用的操作进行判断
	var num int
	fmt.Scan(&num)
	switchType(num)
}

func switchType(n int) {
	switch n {
	case 1:
		// 添加联系人函数
		addPerson()
	case 2:
		deletePerson()
	case 3:
		findPerson()
	case 4:
		editPerson()
	}
}
