package main

import "fmt"

func main() {

	// Initial slice
	s := []int{10, 20, 30}
	fmt.Println("Original Slice:", s)

	// Append
	s = append(s, 40, 50)
	fmt.Println("After Append:", s)

	// Copy
	copySlice := make([]int, len(s))
	copy(copySlice, s)
	fmt.Println("Copied Slice:", copySlice)

	// Remove element (remove element at index 2)
	index := 2
	s = append(s[:index], s[index+1:]...)
	fmt.Println("After Removing element at index 2:", s)
}
