package main

import "fmt"

func main() {
	var a, b int
	var choice int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Result:", a+b)
	case 2:
		fmt.Println("Result:", a-b)
	case 3:
		fmt.Println("Result:", a*b)
	case 4:
		if b != 0 {
			fmt.Println("Result:", a/b)
		} else {
			fmt.Println("Division by zero not allowed")
		}
	default:
		fmt.Println("Invalid Choice")
	}
}
