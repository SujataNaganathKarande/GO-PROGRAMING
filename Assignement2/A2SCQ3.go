package main

import "fmt"

func operations(a int, b int) (int, int, int) {
	sum := a + b
	product := a * b
	diff := a - b
	return sum, product, diff
}

func main() {
	var x, y int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	s, p, d := operations(x, y)

	fmt.Println("Sum:", s)
	fmt.Println("Product:", p)
	fmt.Println("Difference:", d)
}
