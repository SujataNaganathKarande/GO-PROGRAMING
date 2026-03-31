package main

import "fmt"

// function using pointers
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 10
	y := 20

	fmt.Println("Before Swap:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)

	swap(&x, &y)

	fmt.Println("After Swap:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
}
