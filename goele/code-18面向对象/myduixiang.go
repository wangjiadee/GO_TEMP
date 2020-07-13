package main

import "fmt"

type Stroager interface {
	Read()
	Writer()
}
type MDisk struct {
}

// mdisk method
func (m *MDisk) Read() {
	fmt.Println("移动硬盘读取数据")
}
func (m *MDisk) Writer() {
	fmt.Println("移动硬盘写入数据")
}


type UDisk struct {
}
// udisk method

func (u *UDisk) Read() {
	fmt.Println("U盘读取数据")
}
func (u *UDisk) Writer() {
	fmt.Println("U盘写入数据")
}

// 多态定义一个公共的函数 使用接口类型
func Computer(c Stroager){
	c.Read()
	c.Writer()
}


func main() {
	var udisk UDisk
	var mdisk MDisk
	Computer(&udisk)
	Computer(&mdisk)
}