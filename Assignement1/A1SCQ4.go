package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num >= -9 && num <= 9 {
		fmt.Println("It is a Single Digit Number")
	} else {
		fmt.Println("It is Not a Single Digit Number")
	}
}
