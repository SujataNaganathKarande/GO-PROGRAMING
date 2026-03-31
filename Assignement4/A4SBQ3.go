package main

import "fmt"

// Define a structure that holds an array
type ArrayHolder struct {
	data [5]int
}

// Method with pointer receiver to copy elements
func (a *ArrayHolder) copyFrom(source [5]int) {
	for i := 0; i < len(source); i++ {
		a.data[i] = source[i]
	}
}

func main() {
	// Source array
	source := [5]int{10, 20, 30, 40, 50}

	// Destination holder
	var dest ArrayHolder

	// Copy elements using method
	dest.copyFrom(source)

	// Display copied array
	fmt.Println("Source Array     :", source)
	fmt.Println("Destination Array:", dest.data)
}
