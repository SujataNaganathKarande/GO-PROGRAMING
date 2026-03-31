package main

import "fmt"

func main() {
	x := 10
	p := &x  // pointer to x
	pp := &p // pointer to pointer

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)

	fmt.Println("Value stored in p:", p)
	fmt.Println("Value at *p:", *p)

	fmt.Println("Value stored in pp:", pp)
	fmt.Println("Value at *pp:", *pp)
	fmt.Println("Value at **pp:", **pp)
}
