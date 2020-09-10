package main

// Calculator
import "fmt"

// Define Class
type Object struct {
}

//
func (o *Object) GetResult(num1 int, num2 int, op string) int {
	var result int
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	}
	return result
}

// =========================main function============================
func main() {
	var obj Object
	NUM := obj.GetResult(500, 20, "+")
	fmt.Println(NUM)
}

