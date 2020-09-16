/*
 * @Author: Ralph
 * @Type_file[python//GO//json//Yaml//Other]: [GO]
 * @Date: 2020-09-16 10:04:17
 * @LastEditTime: 2020-09-16 10:21:54
 * @FilePath: /go/src/github.com/Ralph.org/File_Operation/main.go
 * @Effect: DO
 */
package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)


func Create_file() {
	f,err := os.Create("/root/go/src/github.com/Ralph.org/File_Operation/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	fmt.Println("Suceessful!")
}

func Open_Write() {
	// f,err := os.Open("/root/go/src/github.com/Ralph.org/File_Operation/test.txt")
	f,err := os.OpenFile("/root/go/src/github.com/Ralph.org/File_Operation/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_,err = f.WriteString("Fucking work")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Writing success!")
}


func Open_Write2()  {
	f, err := os.OpenFile("/root/go/src/github.com/Ralph.org/File_Operation/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer f.Close()
	fmt.Println("successful")

	n, err := f.WriteString("Fucking world\r\n")
	if err != nil {
		fmt.Println("WriteString err:", err)
		return
	}
	fmt.Println("WriteString n = ", n)

	off, _ := f.Seek(-5, io.SeekEnd)
	fmt.Println("off:", off)

	n, _ = f.WriteAt([]byte("1111"), off)
	fmt.Println("WriteAt n :", n)
}

func Read_File()  {
	f, err := os.OpenFile("/root/go/src/github.com/Ralph.org/File_Operation/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer f.Close()
	fmt.Println("successful")

	// 创建一个带有缓冲区(用户缓冲)的 reader
	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n')		// 读一行数据
		if err != nil && err == io.EOF {
			fmt.Println("Read file finish")
			return
		} else if err != nil {
			fmt.Println("ReadBytes err:", err)
		}
		fmt.Print(string(buf))
	}
}

func Copy_File()  {

	// 打开读文件
	f_r, err := os.Open("/root/go/src/github.com/Ralph.org/File_Operation/test.txt")
	if err != nil {
		fmt.Println("Open err: ", err)
		return
	}
	defer f_r.Close()
	// 创建写文件
	f_w, err := os.Create("/root/go/src/github.com/Ralph.org/File_Operation/av.txt")
	if err != nil {
		fmt.Println("Create err: ", err)
		return
	}
defer f_w.Close()

	// 从读文件中获取数据，放到缓冲区中。
	buf := make([]byte, 4096)
	// 循环从读文件中，获取数据，“原封不动的”写到写文件中。
	for {
		n, err := f_r.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Printf("读完。n = %d\n", n)
			return
		}
		f_w.Write(buf[:n])		// 读多少，写多少
	}

}

func Dir_Operation()  {
	// 获取用户输入的目录路径
	fmt.Println("请输入待查询的目录：")
	var path string
	fmt.Scan(&path)

	// 打开目录
	f, err := os.OpenFile(path, os.O_WRONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err: ", err)
		return
	}
	defer f.Close()
	// 读取目录项
	info, err:= f.Readdir(-1)	// -1： 读取目录中所有目录项
	if err != nil {
		fmt.Println("Readdir err: ", err)
		return
	}
	// 变量返回的切片
	for _, fileInfo := range info {
		if fileInfo.IsDir() {			// 是目录
			fmt.Println(fileInfo.Name(), " 是一个目录")
		} else {
			fmt.Println(fileInfo.Name(), " 是一个文件")
		}
	}
}


func main() {
	// Create_file()
	// Open_Write()
	// Open_Write2()
	// Read_File()
	// Copy_File()
	Dir_Operation()
}