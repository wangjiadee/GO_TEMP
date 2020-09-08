package main

import "fmt"

type Robot struct {
	Name  string
	Color string
	Model string
}

// Deep copy
func Deep_Copy() {
	fmt.Println(`深拷贝 内容一样，改变其中一个对象的值时，另一个不会变化。
深拷贝 是复制了数据和内存地址(2个都复制了)`)
	robot1 := Robot{
		Name:  "小白-X型-V1.0",
		Color: "白色",
		Model: "小型",
	}
	robot2 := robot1
	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, &robot2)

	fmt.Println("修改Robot1的Name属性值")
	robot1.Name = "小白-X型-V1.1"

	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, &robot2)

}

// Shallow copy
func Shallow_Copy() {
	fmt.Println("浅拷贝 内容和内存地址一样，改变其中一个对象的值时，另一个同时变化。")
	robot1 := Robot{
		Name:  "小白-X型-V1.0",
		Color: "白色",
		Model: "小型",
	}
	robot2 := &robot1
	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, robot2)

	fmt.Println("在这里面修改Robot1的Name和Color属性")
	robot1.Name = "小黑-X型-V1.1"
	robot1.Color = "黑色"

	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, robot2)

}

// new copy ---> shallow copy
func New_Copy() {
	fmt.Println("浅拷贝 使用new方式")
	robot1 := new(Robot)
	robot1.Name = "小白-X型-V1.0"
	robot1.Color = "白色"
	robot1.Model = "小型"

	robot2 := robot1
	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, robot1)
	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, robot2)

	fmt.Println("在这里面修改Robot1的Name和Color属性")
	robot1.Name = "小蓝-X型-V1.2"
	robot1.Color = "蓝色"

	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, robot1)
	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, robot2)
}

/////////////////////start main function
func main() {
	// Deep_Copy()
	// Shallow_Copy()
	New_Copy()
}
