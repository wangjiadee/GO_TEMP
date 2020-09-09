package main

import "fmt"

// 定义类
type MDisk struct {
}

type UDisk struct {
}

// 定义类的方法
func (m *MDisk) Read() {
	fmt.Println("移动硬盘读取数据")
}
func (m *MDisk) Writer() {
	fmt.Println("移动硬盘写入数据")
}
func (u *UDisk) Read() {
	fmt.Println("U盘读取数据")
}
func (u *UDisk) Writer() {
	fmt.Println("U盘写入数据")
}

// 定义接口
type Stroager interface {
	Read()
	Writer()
}

// 定义多态的函数
func Computer(c Stroager) {
	c.Read()
	c.Writer()
}

// =========================main function============================
func main() {
	var udisk UDisk
	var mdisk MDisk
	Computer(&udisk)
	Computer(&mdisk)
}

