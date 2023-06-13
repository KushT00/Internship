package main

import (
	"fmt"
)

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func sub(num1, num2 float64) float64 {
	return num2 - num1
}

func mul(num1, num2 float64) float64 {
	return num1 * num2
}

func div(num1, num2 float64) float64 {
	if num2 != 0 {
		return num1 / num2
	}

	fmt.Printf("Cannot be divided by 0 ")
	return 0
}
func calc() {
	var a, b, op float64

	fmt.Println("Enter 1st number: ")
	fmt.Scanln(&a)

	fmt.Println("Enter 2nd number: ")
	fmt.Scanln(&b)

	fmt.Println("Enter Operation you want to perform ")
	fmt.Println("1 Addition ")
	fmt.Println("2 Subtraction")
	fmt.Println("3 Multiplication ")
	fmt.Println("4 division")
	fmt.Scanln(&op)

	switch op {
	case 1:
		ans := add(a, b)
		fmt.Println("Result: ", ans)

	case 2:
		ans := sub(a, b)
		fmt.Println("Result: ", ans)

	case 3:
		ans := mul(a, b)
		fmt.Println("Result: ", ans)

	case 4:
		ans := div(a, b)
		fmt.Println("Result: ", ans)

	default:
		fmt.Println("Invalid operation number!")

	}
}

func main() {
	calc()
}
