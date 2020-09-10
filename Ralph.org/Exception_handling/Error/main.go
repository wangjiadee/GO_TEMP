package main

import (
	"errors"
	"fmt"

)

func TestError(num1 int, num2 int) (result int, return_err error) {
	return_err = nil
	if num2 == 0 {
		return_err = errors.New("ERROR: 除数不能为0!")
		return
	}
	result = num1 / num2
	return
}

// =======================================Main Function start==========================
func main() {
	num, return_err := TestError(10, 0)
	if return_err != nil {
		fmt.Println(return_err)
	} else {
		fmt.Println(num)
	}

}
