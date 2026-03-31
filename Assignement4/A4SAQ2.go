package main

import "fmt"

type Calculator struct{}

func (c Calculator) Multiply(a int, b int) int {
	return a * b
}

func main() {
	calc := Calculator{}
	result := calc.Multiply(8, 12)
	fmt.Printf("Multiplication: %d\n", result)
}
