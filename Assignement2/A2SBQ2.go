package main

import "fmt"

func calculate(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return // automatically returns sum and diff
}

func main() {
	var x, y int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	s, d := calculate(x, y)

	fmt.Println("Sum:", s)
	fmt.Println("Difference:", d)
}
