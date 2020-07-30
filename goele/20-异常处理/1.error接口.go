/*
 * @Author: Ralph
 * @Type_file: GO
 * @Date: 2020-04-27 21:59:02
 * @LastEditTime: 2020-07-30 21:39:20
 * @FilePath: \go\goele\20-异常处理\1.error接口.go
 * @Effect: error异常处理
 */

/*
类似python中的try

*/

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
