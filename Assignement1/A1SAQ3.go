package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println("Before Swap:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	a, b = b, a

	fmt.Println("After Swap:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
