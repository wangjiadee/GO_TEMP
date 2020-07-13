package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	file, err := os.Open("c:/Test/file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buffer := make([]byte, 10)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF { // 表示达到文件末尾了。
			break
		}

		fmt.Print(string(buffer[:n]))
	}

}
