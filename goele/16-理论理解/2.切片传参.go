package main

	import "fmt"
	//切片底层就是一个结构体
	func main() {
		s := make([]int, 5, 5)
		Modify(s)
		fmt.Println(s) //[0 0 0 0 0]
	}
	func Modify(sli []int) {
		for i := 0; i < 5; i++ {
			// sli[i] = i
			sli = append(sli, i)
		}
		fmt.Println("sli:",sli) //sli: [0 0 0 0 0 0 1 2 3 4]

	}
