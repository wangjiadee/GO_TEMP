package main

import (
	"os"
	"fmt"
)

func main()  {
	// 创建文件，需要指定文件的存储路径以及文件名称，注意“/”
	file,err:=os.Create("C:/Test/a.txt")
	//判断是否出现异常
	if err!=nil{
		fmt.Println(err)
		//file.Close()
	}
	defer file.Close()
	//对创建的文件进行相关的操作。

	//关闭
	//file.Close()


}
