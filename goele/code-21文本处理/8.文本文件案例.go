package main

import (
	"os"
	"fmt"
	"io"
)

//文件拷贝，将已有的文件复制一份，同时重新命名。
func main()  {
  // 1: 打开原有文件(file.txt)
  file1,err:=os.Open("c:/Test/file.txt")
  if err!=nil{
  	fmt.Println(err)
  }
  // 2: 创建一个新的文件
  file2,err:=os.Create("c:/Test/file2.txt")
  if err!=nil{
  	fmt.Println(err)
  }
  defer file1.Close()
  defer  file2.Close()
  // 3: 将原有文件中的内容读取出来，然后写入到新的文件中。
  buffer:=make([]byte,1024*2)
	for  {
		n,err:=file1.Read(buffer)
		if err==io.EOF {
			break
		}
		file2.Write(buffer[:n])
	}

  // 4: 关闭文件
}
