package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter number of terms: ")
	fmt.Scan(&n)

	a := 0
	b := 1

	fmt.Println("Fibonacci Series:")

	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")
		next := a + b
		a = b
		b = next
	}
}
