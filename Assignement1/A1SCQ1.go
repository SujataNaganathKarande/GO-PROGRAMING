package main

import "fmt"

func main() {
	str1 := "Hello "
	str2 := "Ram"

	p1 := &str1
	p2 := &str2

	result := *p1 + *p2

	fmt.Println("Concatenated String:", result)
}
