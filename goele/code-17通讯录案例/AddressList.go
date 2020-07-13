package main

import (
	"fmt"
)

type Person struct {
	userName     string
	addressPhone map[string]string // key :表示出电话的类型，value:电话
}

var personList = make([]Person, 0)
// 添加联系人
func addPerson() {
	// 1: 定义结构体表示联系人的信息
	// 2: 定义切片保存多个人的联系信息。
	// 3； 向切片中保存数据。
	//  3.1 添加姓名
	var name string
	var address string
	var phone string
	var exit string                            // 表示退出电话的录入
	var addressPhone = make(map[string]string) // 保存电话的类型和电话，电话类型作为key
	fmt.Println("请输入姓名")
	fmt.Scan(&name)

	for {
		//  3.2 保存电话类型
		fmt.Println("请输入电话类型")
		fmt.Scan(&address)
		//  3.3 保存电话号码
		fmt.Println("请输入电话号码")
		fmt.Scan(&phone)
		// 将电话以及电话类型存储到addressPhone 中。
		addressPhone[address] = phone
		fmt.Println("如果结束电话的录入，请按Q")
		fmt.Scan(&exit)
		if exit == "Q" {
			break;
		} else {
			continue
		}

	} //将联系人的信息存储到切片中。
	personList = append(personList, Person{userName: name, addressPhone: addressPhone})
	//fmt.Println(personList)
	showPersonList() // 调用函数展示联系人的信息
}

// 展示切片中存储的联系人信息。
func showPersonList() {
	// 1: 判断一下切片中是否有数据。
	if len(personList) == 0 {
		fmt.Println("暂时没有联系人信息")
	} else {
		// 2；可以通过循环的方式打印切片中的数据。
		for _, value := range personList {
			fmt.Println("姓名：", value.userName)
			for k, v := range value.addressPhone {
				fmt.Println("电话类型：", k)
				fmt.Println("电话号码：", v)
			}

		}
	}

}

// 删除联系人信息
func removePerson() {
	// 1: 输入要删除的联系人的姓名
	var name string
	var index int = -1 // 记录要删除的联系人信息在切片中的下标。
	fmt.Println("请输入要删除的联系人姓名：")
	fmt.Scan(&name)
	// 2: 判断切片中是否存储了要删除的联系人信息。
	for i := 0; i < len(personList); i++ {
		if personList[i].userName == name {
			index = i
			break
		}
	}
	// 3: 记录要删除的联系人信息在切片中的位置（下标）
	// 4： 删除操作
	// index = 3  //   5,6,7,9,10,11,12
	//8,
	if index != -1 {
		personList = append(personList[0:index], personList[index+1:]...) // append函数第二个参数如果是切片后面要给三个点。
	}
	showPersonList()
}

// 查找联系人信息
func findPerson() *Person {
	// 1: 输入要查询的联系人姓名
	var name string
	var index int = -1 // 记录找到的联系人信息在切片中的下标
	fmt.Println("请输入要查询的联系人姓名：")
	fmt.Scan(&name)
	// 2: 根据输入的联系人姓名，查找对应的联系信息
	for k, value := range personList {
		if value.userName == name {
			index = k
			fmt.Println("联系人姓名：", value.userName)
			for key, v := range value.addressPhone {
				fmt.Printf("%s:%s\n", key, v)
			}
		}
	}
	// 3: 打印输出结果
	if index == -1 {
		fmt.Println("没有找到联系人信息")
		return nil

	} else {
		return &personList[index]
	}

}

// 编辑联系人信息
func editPerson() {
	// 1: 查找到要编辑的联系人信息
	var name string // 存储新的联系人姓名
	var p *Person
	var num int                  // 存储修改的数据的类型
	var menu = make([]string, 0) // 保存电话类型，方便后面修改
	var pNum int                 // 编辑的电话类型编号
	var phone string             // 新的电话号码
	p = findPerson()
	if p != nil {
		for {
			fmt.Println("编辑用户名称请按:5,编辑电话请按:6,退出请按:7")
			fmt.Scan(&num)
			switch num {
			case 5:
				fmt.Println("请输入新的姓名：")
				fmt.Scan(&name)
				p.userName = name
				showPersonList()
			case 6:
				// 编辑联系电话
				// 1: 展示联系人所有的电话信息
				var j int
				for key, value := range p.addressPhone {
					fmt.Println("编辑(", key, ")", value, "请按：", j)
					menu = append(menu, key)
					j++
				}
				// menu 公司，学校
				//       0    1
				fmt.Println("请输入编辑号码的类型：")
				fmt.Scan(&pNum) //1
				// 2: 完成修改
				for index, v := range menu {
					if index == pNum {
						fmt.Println("请输入新的电话号码：")
						fmt.Scan(&phone)
						p.addressPhone[v] = phone
					}
				}
			}
			if num == 7 {
				break
			}

		}

	} else {
		fmt.Println("没有找到要编辑的联系人信息")
	}
	// 2: 进行编辑
}

func main() {
	for {
		scanNum()
	}

}
func scanNum() {
	// 1: 给出相应的操作提示
	fmt.Println("添加联系人信息，请按1")
	fmt.Println("删除联系人信息，请按2")
	fmt.Println("查询联系人信息，请按3")
	fmt.Println("编辑联系人信息，请按4")
	// 2: 对用户输入的数字进行判断。
	var num int // 保存用户输入的数字
	fmt.Scan(&num)
	switchType(num)
}

// 对输入的内容进行判断，决定执行哪块操作
func switchType(n int) {
	switch n {
	case 1:
		// 添加联系人的操作
		addPerson()
	case 2:
		// 删除联系人的操作
		removePerson()
	case 3:
		// 查询联系人的操作
		findPerson()
	case 4:
		// 编辑联系人的操作
		editPerson()
	}
}
