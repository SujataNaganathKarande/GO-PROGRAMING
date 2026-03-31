package main

import "fmt"

// function to add two numbers
func add(a int, b int) int {
	return a + b
}

func main() {
	var num1, num2 int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&num1, &num2)

	result := add(num1, num2)

	fmt.Println("Addition is:", result)
}
