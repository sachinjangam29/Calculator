package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multiple(a, b int) int {
	return a * b
}

func division(a, b int) int {
	if a == 0 || b == 0 {
		//fmt.Println("Enter the valid numbers and not /'0/' ")
		return -1
	}
	return (a / b)
}

func main() {
	fmt.Println("Welcome to the Calculator screen")

	result := add(10, 5)
	fmt.Println("The addition of the number is ", result)

	result = sub(10, 5)
	fmt.Println("The subtraction of the number is ", result)

	result = multiple(10, 5)
	fmt.Println("The multiplication of the number is ", result)

	result = division(10, 5)
	fmt.Println("The division of the number is ", result)

}
