package main

import "fmt"

func changeValue(num int) {
	num = 100
	fmt.Println("Inside function:", num)
}

func main() {
	x := 50

	fmt.Println("Before function call:", x)

	changeValue(x)

	fmt.Println("After function call:", x)
}
