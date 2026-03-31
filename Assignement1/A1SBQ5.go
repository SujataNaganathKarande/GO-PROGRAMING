package main

import "fmt"

func main() {
	ptr := new(int) // allocate memory for int

	fmt.Println("Default value:", *ptr)

	*ptr = 25

	fmt.Println("Updated value:", *ptr)
	fmt.Println("Address:", ptr)
}
