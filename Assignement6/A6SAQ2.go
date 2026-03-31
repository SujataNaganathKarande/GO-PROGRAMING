package main

import (
	"Assignement6/calculator"
	"fmt"
)

func main() {

	var a, b, choice int

	fmt.Println("Enter two numbers:")
	fmt.Scan(&a, &b)

	fmt.Println("1.Addition")
	fmt.Println("2.Subtraction")
	fmt.Println("3.Multiplication")
	fmt.Println("4.Division")

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		fmt.Println("Result:", calculator.Add(a, b))

	case 2:
		fmt.Println("Result:", calculator.Sub(a, b))

	case 3:
		fmt.Println("Result:", calculator.Mul(a, b))

	case 4:
		fmt.Println("Result:", calculator.Div(a, b))

	default:
		fmt.Println("Invalid choice")
	}
}
