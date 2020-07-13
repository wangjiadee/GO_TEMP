package main

import "fmt"

type Stroager interface {
	Read()
	Writer()
}

// 移动硬盘
type MDisk struct {
}

func (m *MDisk) Read() {
	fmt.Println("移动硬盘读取数据")
}
func (m *MDisk) Writer() {
	fmt.Println("移动硬盘写入数据")
}

// U盘
type UDisk struct {
}

func (u *UDisk) Read() {
	fmt.Println("U盘读取数据")
}
func (u *UDisk) Writer() {
	fmt.Println("U盘写入数据")
}
// 定义一个函数
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
